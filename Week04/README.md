Question
---
按照自己的构想，写一个项目满足基本的目录结构和工程，代码需要包含对数据层、业务层、API 注册，以及 main 函数对于服务的注册和启动，信号处理，使用 Wire 构建依赖。可以使用自己熟悉的框架。

Answer
---
基本结构:
```
├── README.md
├── api             # rpc文件
├── cmd
│   └── demo        #demo 项目
├── go.mod
├── go.sum
├── configs         # 配置文件
├── internal
│   ├── biz         # 业务组装层
│   ├── data        # 数据层
│   ├── pkg         # 项目共有库
│   └── service     # API实现层
```

当前存在问题:
- 并没有用到wire .wire-demo 实现了个wire 的样例，（demo中目前觉得没有必要或可用的地方）
- todo 日志组件、信号处理等
