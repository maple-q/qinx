package qiface

type IServer interface {
    Start()

    Stop()

    Serve()
    // 路由功能：给当前服务注册一个路由业务方法，供客户端连接使用
    AddRouter(router IRouter)
}
