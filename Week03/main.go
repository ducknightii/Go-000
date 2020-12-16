package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

type httpServer struct {
	g      *errgroup.Group
	ctx    context.Context
	cancel func()
}

func main() {
	h := Init()
	h.AddServer(":18001", nil)
	h.AddServer(":18002", nil)
	h.AddServer(":18003", nil)
	h.Wait()
}

func Init() *httpServer {
	var h httpServer
	var ctx context.Context
	ctx, h.cancel = context.WithCancel(context.Background())
	h.g, h.ctx = errgroup.WithContext(ctx)
	return &h
}

func (h *httpServer) AddServer(addr string, handler http.Handler) {
	h.g.Go(func() (err error) {
		defer func() {
			if r := recover(); r != nil {
				buf := make([]byte, 64<<10)
				buf = buf[:runtime.Stack(buf, false)]
				// err 赋值返回 return 无法接收
				err = fmt.Errorf("errgroup: panic recovered: %s\n%s", r, buf)
			}
		}()

		s := http.Server{
			Addr:    addr,
			Handler: handler,
		}
		go func() {
			<-h.ctx.Done()
			fmt.Println("Addr: ", addr, " Closed!!")
			ctx, _ := context.WithTimeout(context.Background(), time.Second)
			s.Shutdown(ctx)
		}()
		return s.ListenAndServe()
	})
}

func (h *httpServer) Wait() {
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGUSR1, syscall.SIGUSR2)
		s := <-c
		fmt.Println("Exit Single: ", s)
		h.cancel()
	}()

	if err := h.g.Wait(); err != nil {
		fmt.Println("Err Closed: ", err)
	}
}
