package main

import (
    "fmt"
    "net"
    "time"
)


func main() {
    fmt.Println("Client Test ... start")
    time.Sleep(3 * time.Second)

    conn, err := net.Dial("tcp", "127.0.0.1:7777")
    if err != nil {
        fmt.Println("start client err: ", err)
        return
    }

    for {
        _, err = conn.Write([]byte("hahaha"))
        if err != nil {
            fmt.Println("write error: ", err)
            return
        }

        buf := make([]byte, 512)
        cnt, err := conn.Read(buf)
        if err != nil {
            fmt.Println("read from server err: ", err)
            return
        }

        fmt.Printf("server call back %s, cnt = %d\n", buf, cnt)
        time.Sleep(1 * time.Second)
    }
}
