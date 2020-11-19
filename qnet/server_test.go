package qnet

import (
    "fmt"
    "net"
    "testing"
    "time"
)


// 模拟客户端
func ClientTest() {
    fmt.Println("Client Test ... Start")

    time.Sleep(3 * time.Second)

    conn, err := net.Dial("tcp", "127.0.0.1:7777")
    if err != nil {
        fmt.Println("client start err: ", err)
        return
    }

    for {
        _, err := conn.Write([]byte("hello ZINX"))
        if err != nil {
            fmt.Println("write error err: ", err)
            return
        }

        buf := make([]byte, 512)
        cnt, err := conn.Read(buf)
        if err != nil {
            fmt.Println("read buf err: ", err)
            return
        }

        fmt.Printf("Server call back:%s, cnt = %d\n", buf, cnt)
        time.Sleep(1 * time.Second)
    }
}


// 测试Server模块
func TestServer(t *testing.T) {
    s := NewServer("[qinx v0.1]")

    go ClientTest()

    s.Serve()
}
