Question
---
用 Go 实现一个 tcp server ，用两个 goroutine 读写 conn，两个 goroutine 通过 chan 可以传递 message，能够正确退出

Answer
---
- 实现了对于接入的客户端，选取其中一个空闲连接，接收其消息
- 当消费的客户端断开连接后，从其余空闲连接中选取一个，重新进行选取
- 当只有两个客户端时，也就实现了两者相互通话的效果

TODO
- 两两配对相互通信/广播
