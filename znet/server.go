package znet

import (
	"fmt"
	"net"
)

// Server 实现IServer接口的服务类
type Server struct {
	Name      string // 服务器名称
	IPVersion string // tcp，协议名称
	IP        string // ip地址
	Port      int    // 绑定服务的端口
}

// NewServer 创建一个服务
func NewServer(name string) *Server {
	srv := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8000,
	}

	return srv
}

// Start 启动服务
func (s *Server) Start() {
	// 整个Start过程启动一个协程来完成，不占用使用方主协程
	go func() {

		// 1. 通过 ResolveTCPAddr 方法获取 *net.TCPAddr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err: ", err)
			return
		}

		// 2. 使用 ListenTCP 方法监听服务端口
		lis, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen to failed: ", err)
			return
		}

		fmt.Printf("start zinx %s server on %s:%d\n", s.Name, s.IP, s.Port)

		// 3. 等待接受网络连接
		for {
			conn, err := lis.AcceptTCP()
			if err != nil {
				fmt.Println("accept connect failed: ", err)
				continue
			}

			// 每一个连接单独开启一个协程处理业务
			go func() {
				// 不断从客户端循环读取数据
				for {
					buf := make([]byte, 1024)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("read data failed: ", err)
						continue
					}
					// 做一个简答，读到什么就回显什么
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write data failed: ", err)
						continue
					}
				}
			}()
		}
	}()
}

// Stop 停止服务
func (s *Server) Stop() {
	// TODO
}

// Serve 启动服务处理函数
func (s *Server) Serve() {
	s.Start()

	// TODO Server.Serve() 是否在启动服务的时候 还要处理其他的事情呢 可以在这里添加
}
