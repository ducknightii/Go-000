package demo

import (
	"fmt"
	"time"
)

// 	Never start a goroutine without knowning when it will stop
// ch 没有关闭 goroutine 泄漏 直到主线程关闭
func leak() {
	ch := make(chan int)
	go func() {
		val := <-ch
		fmt.Println("Val: ", val)
	}()
	//close(ch)
	time.Sleep(time.Second)
}
