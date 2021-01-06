package main

import (
	"fmt"
	"sync"
	"time"
)

type metric struct {
	success int
	fail    int
}

type bucket struct {
	m metric
	t time.Time
}

type SlideWindowCounter struct {
	buckets  []*bucket
	size     int
	interval time.Duration // bucket的间隙(s)
	head     int
	tail     int
	mux      sync.RWMutex
}

func NewSlideWindowCounter(size int, interval time.Duration) *SlideWindowCounter {
	return &SlideWindowCounter{
		size:     size,
		interval: interval,
		buckets:  make([]*bucket, size),
	}
}

func (s *SlideWindowCounter) Incr(successNum, failNum int) {
	s.mux.Lock()
	defer s.mux.Unlock()

	current := time.Now()
	// 还是空的
	if s.head == 0 && s.buckets[s.head] == nil {
		s.buckets[s.head] = &bucket{
			m: metric{
				success: successNum,
				fail:    failNum,
			},
			t: current,
		}
		return
	}

	// 还在tail 区间内
	if current.Sub(s.buckets[s.tail].t) < s.interval*time.Second {
		s.buckets[s.tail].m.success += successNum
		s.buckets[s.tail].m.fail += failNum
		return
	}
	// 如果 current-tail > size * interval 存储的数据无用 可以清空队列  重新初始化了
	if current.Sub(s.buckets[s.tail].t) > s.interval*time.Second*time.Duration(s.size) {
		s.buckets = make([]*bucket, s.size)
		s.head = 0
		s.tail = 0
		s.buckets[s.head] = &bucket{
			m: metric{
				success: successNum,
				fail:    failNum,
			},
			t: current,
		}
		return
	}
	// 从tail 开始 遍历找到可以插入的bucket
	_t := s.buckets[s.tail].t
	for current.Sub(s.buckets[s.tail].t) >= s.interval*time.Second {
		s.tail = (s.tail + 1) % s.size
		_t = _t.Add(s.interval * time.Second)
		s.buckets[s.tail] = &bucket{
			m: metric{
				success: 0,
				fail:    0,
			},
			t: _t,
		}
		if s.tail == s.head {
			s.head = (s.head + 1) % s.size
		}
	}
	s.buckets[s.tail].m.success += successNum
	s.buckets[s.tail].m.fail += failNum
	return
}

func (s *SlideWindowCounter) Counter() (successNum, failNum int) {
	s.mux.RLock()
	defer s.mux.RUnlock()

	for i := s.head; i < s.head+s.size && s.buckets[i%s.size] != nil; i++ {
		successNum += s.buckets[i%s.size].m.success
		failNum += s.buckets[i%s.size].m.fail
	}
	return
}

func main() {
	s := NewSlideWindowCounter(10, 1)
	for i := 0; i < 100; i++ {
		s.Incr(i, 100-i)
		time.Sleep(time.Second)
		success, fail := s.Counter()
		fmt.Println(i, "==>", success, fail)
	}
}
