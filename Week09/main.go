package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

type Ch struct {
	used bool
	ch   chan []byte
}

// 存储所有在线连接
type Channels struct {
	m   sync.Mutex
	chs map[int]*Ch // 每个链接一个ch
}

var channels Channels

func main() {
	listen, err := net.Listen("tcp", ":10000")

	if err != nil {
		log.Fatalf("listen error: %v\n", err)
	}

	channels.chs = make(map[int]*Ch)

	id := 0

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("accept error: %v\n", err)
			continue
		}
		id++
		ch := make(chan []byte, 10)
		ctx, cancel := context.WithCancel(context.Background())

		log.Printf("accept: %d\n", id)

		channels.m.Lock()

		go handleRead(conn, id, ch, cancel)
		go handleWrite(ctx, conn, id)

		channels.chs[id] = &Ch{
			used: false,
			ch:   ch,
		}

		channels.m.Unlock()
	}
}

func handleRead(conn net.Conn, id int, ch chan<- []byte, cancelFunc context.CancelFunc) {
	defer func() {
		close(ch)
		channels.m.Lock()
		delete(channels.chs, id)
		channels.m.Unlock()
		cancelFunc()
		conn.Close()
	}()

	rd := bufio.NewReader(conn)
	for {
		line, _, err := rd.ReadLine()
		if err != nil {
			log.Printf("%d: read error:%v\n", id, err)
			return
		}
		ch <- line
	}
}

func handleWrite(ctx context.Context, conn net.Conn, id int) {
	defer func() {
		channels.m.Lock()
		delete(channels.chs, id)
		channels.m.Unlock()
	}()

	wr := bufio.NewWriter(conn)
	wr.WriteString(fmt.Sprintf("Your ID: %d\n", id))
	wr.Flush()

	// todo 退出检查
	fromID, ch := pickSender(id)

	for {
		select {
		case line, ok := <-ch:
			if !ok {
				wr.WriteString(fmt.Sprintf("%d leaved\n", fromID))
				wr.Flush()
				// 重新选
				fromID, ch = pickSender(id)
				break
			}
			wr.WriteString(fmt.Sprintf("%d: ", fromID))
			wr.Write(line)
			wr.WriteString("\n")
			err := wr.Flush()
			if err != nil {
				log.Printf("%d: write error:%v\n", id, err)
				return
			}
		case <-ctx.Done():
			// 释放消费的ch
			channels.m.Lock()
			channels.chs[fromID].used = false
			channels.m.Unlock()
			return
		}
	}
}

// 选择一个对话
func pickSender(id int) (int, <-chan []byte) {
	// 选一个 非自己的chan
	var ch <-chan []byte
	var fromID int
	for {
		time.Sleep(time.Millisecond * 100)
		channels.m.Lock()
		if len(channels.chs) < 2 {
			channels.m.Unlock()
			continue
		}
		for _id, _u := range channels.chs {
			if _id != id && !_u.used {
				_u.used = true
				ch = _u.ch
				channels.m.Unlock()
				fromID = _id
				log.Printf("%d --> %d", _id, id)
				break
			}
		}
		if ch != nil {
			break
		}
	}

	return fromID, ch
}
