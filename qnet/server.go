package qnet


import "qinx/qiface"
import "fmt"
import "time"
import "net"
import "errors"


type Server struct {
    Name string
    IPVersion string
    IP string
    Port int
}


// 定义当前客户端连接的handle api
func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error{
    // 回显业务
    fmt.Println("[Conn Handle] CallBackToClient...")
    if _, err := conn.Write(data[:cnt]); err != nil {
        fmt.Println("write back buff err: ", err)
        return errors.New("CallBackToClient error")
    }

    return nil
}


func NewServer(name string) qiface.IServer {
    s := &Server {
        Name: name,
        IPVersion: "tcp4",
        IP: "0.0.0.0",
        Port: 7777,
    }
    return s
}


func (s *Server) Start() {
    fmt.Printf("[START] Server listening at IP: %s, Port: %d\n", s.IP, s.Port)

    go func() {
        addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
        if err != nil {
            fmt.Println("resolve tcp add err: ", err)
            return
        }

        // 监听
        listener, err := net.ListenTCP(s.IPVersion, addr)
        if err != nil {
            fmt.Println("listen ", s.IPVersion, " err: ", err)
            return
        }

        fmt.Println("start Qinx server ", s.Name, " success, now listening...")
        // TODO 应该有一个自动生成ID的方法
        var cid uint32
        cid = 0

        // 等待连接
        for {
            conn, err := listener.AcceptTCP()
            if err != nil {
                fmt.Println("Accept err ", err)
                continue
            }

            // TODO 设置服务器最大连接数
            dealConn := NewConnection(conn, cid, CallBackToClient)
            cid ++
            // 启动当前连接的处理业务，这里每个请求过来都是相同的处理逻辑
            go dealConn.Start()
        }
    }()
}


func (s *Server) Stop() {
    fmt.Println("[STOP] Qinx server, name ", s.Name)
    // TODO 清理连接信息
}


func (s *Server) Serve() {
    s.Start()

    // 主goroutine等待
    for {
        time.Sleep(10 * time.Second)
    }
}
