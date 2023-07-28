package ziface

import "net"

// IConnection 定义连接接口
type IConnection interface {
	Start()            // 启动连接，让连接开始工作
	Stop()             // 结束连接
	GetConnID() uint32 // 获取远程客户端地址信息 RemoteAddr() net.Addr
	GetTCPConnection() *net.TCPConn
}

// HandFunc 定义统一处理链接业务的函数类型
type HandFunc func(*net.TCPConn, []byte, int) error
