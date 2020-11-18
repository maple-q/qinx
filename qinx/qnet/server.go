package qnet


import "qinx/qiface"
import "fmt"
import "time"
import "net"


type Server struct {
    Name string
    IPVersion string
    IP string
    Port int
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

        // 等待连接
        for {
            conn, err := listener.AcceptTCP()
            if err != nil {
                fmt.Println("Accept err ", err)
                continue
            }

            // TODO 设置服务器最大连接数
            // TODO 此处应该有handler和conn绑定
            go func() {
                for {
                    buf := make([]byte, 512)
                    cnt, err := conn.Read(buf)
                    if err != nil {
                        fmt.Println("recv buf err ", err)
                        continue
                    }

                    // 回显
                    if _, err := conn.Write(buf[:cnt]); err != nil {
                        fmt.Println("write back buf err ", err)
                        continue
                    }
                }
            }()
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
