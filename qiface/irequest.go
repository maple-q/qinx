package qiface


// IRequest接口，将客户端请求的连接信息和请求的数据包装到了Request里面
type IRequest interface {
    // 获取请求连接对象
    GetConnection() IConnection
    // 获取请求信息数据
    GetData() []byte
}
