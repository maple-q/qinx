package main

import (
    "qinx/qnet"
    "qinx/qiface"
    "fmt"
)

// ping test 自定义路由
type PingRouter struct {
    qnet.BaseRouter
}

// Test PreHandle
func (this *PingRouter) PreHandle(request qiface.IRequest) {
    fmt.Println("Call Router PreHandle")
    _, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping...\n"))
    if err != nil {
        fmt.Println("call back ping error: ", err)
    }
}


// Test Handle
func (this *PingRouter) Handle(request qiface.IRequest) {
    fmt.Println("Call PingRouter Handl")
    _, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...ping..ping\n"))
    if err != nil {
        fmt.Println("call back ping ping ping error")
    }
}


// Test PostHandle
func (this *PingRouter) PostHandle(request qiface.IRequest) {
    fmt.Println("Call Router PostHandle")
    _, err := request.GetConnection().GetTCPConnection().Write([]byte("After...ping\n"))
    if err != nil {
        fmt.Println("call back ping ping ping error: ", err)
    }

}


func main() {
    s := qnet.NewServer("[qinx v0.1]")
    s.AddRouter(&PingRouter{})
    s.Serve()
}
