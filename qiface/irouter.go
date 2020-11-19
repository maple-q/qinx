package qiface


// 路由接口，这里路由是使用框架的用户给该连接定义的处理业务方法
type IRouter interface {
    PreHandle(request IRequest)
    Handle(request IRequest)
    PostHandle(request IRequest)
}
