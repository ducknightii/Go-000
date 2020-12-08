Question:
---
基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够 一个退出，全部注销退出。

Answer:
---
构造结构体
```
type httpServer struct {
	g   *errgroup.Group
	wg  sync.WaitGroup   // 并行计数
	ctx context.Context  // 每个goroutine中监听状态 优雅退出
	cancel func()  // 全局close
}

func (h *httpServer)Init(){} // 初始化
func (h *httpServer) AddServer(addr string, handler http.Handler) {} // 注册监听端口
func (h *httpServer) Wait() {} // 监听服务以及关闭信号
```

