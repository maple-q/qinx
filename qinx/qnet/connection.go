package qnet

import (
    "fmt"
    "net"
    "qinx/qiface"
)


type Connection struct {
    // 当前连接的socket TCP套接字
    Conn *net.TCPConn
    // 当前连接的ID
    ConnID uint32
    // 当前连接状态
    isClosed bool
    // 该连接的处理方法api
    handleAPI qiface.HandFunc
    // 告知该连接已经退出的channel
    ExitBuffChan chan bool
}


// 创建连接的方法
func NewConnection(conn *net.TCPConn, connID uint32, callbackAPI qiface.HandFunc) *Connection {
    return &Connection {
        Conn:        conn,
        ConnID:      connID,
        isClosed:    false,
        handleAPI:   callbackAPI,
        ExitBuffChan:make(chan bool, 1),
    }
}

// 从当前连接获取原始的socket TCPConn
func (c *Connection) GetTCPConnection() *net.TCPConn {
    return c.Conn
}


// 获取连接ID
func (c *Connection) GetConnID() uint32 {
    return c.ConnID
}


// 获取远程客户端的地址
func (c *Connection) RemoteAddr() net.Addr {
    return c.Conn.RemoteAddr()
}


// 处理连接读取数据的goroutine
func (c *Connection) StartReader() {
    fmt.Println("Reader Goroutine is running")
    defer fmt.Println(c.RemoteAddr().String(), " conn reader exit!")
    defer c.Stop()

    for {
        buf := make([]buf, 512)
        cnt, err := c.Conn.Read(buf)
        if err != nil {
            fmt.Println("recv buf err: ", err)
            c.ExitBuffChan <- true
            continue
        }

        // 调用当前连接的业务
        if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
            fmt.Println("connID ", c.ConnID, " handle is error")
            c.ExitBuffChan <- true
            return
        }
    }
}


// 启动连接，让当前连接开始工作
func (c *Connection) Start() {
    go c.StartReader()

    for {
        select {
        case <- c.ExitBuffChan:
            return
        }
    }
}


// 停止连接
func (c *Connection) Stop() {
    // 如果当前连接已经关闭，则直接返回
    if c.isClosed == true {
        return
    }

    c.isClosed = true

    // 关闭连接，发送消息到chan
    c.Conn.Close()
    c.ExitBuffChan <- true
    close(c.ExitBuffChan)
}
