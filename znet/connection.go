package znet

import (
	"fmt"
	"net"

	"github.com/lightsaid/zinx/ziface"
)

// Connection 定义类型链接结构体
// 当tcp accecpt 链接后,交由Connection对象处理业务
type Connection struct {
	Conn         *net.TCPConn    // 当前连接的 socket tcp 套接字
	ConnID       uint32          // 每一个链接分配唯一ID（sessionId）
	isClosed     bool            // 当前链接关闭状态
	handleAPI    ziface.HandFunc // 该链接的处理方法api
	ExitBuffChan chan bool       // 告知该链接已经退出/停止channel
}

// NewConnection 创建链接处理对象的方法
func NewConnection(conn *net.TCPConn, connID uint32, callbackApi ziface.HandFunc) *Connection {
	c := &Connection{
		Conn:         conn,
		ConnID:       connID,
		isClosed:     false,
		handleAPI:    callbackApi,
		ExitBuffChan: make(chan bool),
	}

	return c
}

// StartReader 链接的读方法
// 当创建链接成功后，此方法就会启动一直监听等待对方写入数据
// 具体什么时候退出呢？应该由调用者判断，进而停止goroutine,停止for循环
func (c *Connection) StartReader() {
	fmt.Println("reader goroutine is running")
	defer fmt.Println(c.Conn.RemoteAddr().String(), " conn reader exit!")
	defer c.Stop()

	for {
		buf := make([]byte, 1024)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("read buf err: ", err)
			c.ExitBuffChan <- true
			continue
		}

		// 调用当前链接业务的处理方法
		if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
			fmt.Println("connID: ", c.ConnID, " handle is error")
			c.ExitBuffChan <- true
			return
		}

	}
}

// Start 启动链接，让当前链接开始工作
func (c *Connection) Start() {
	// 开启处理该链接的读工作
	go c.StartReader()

	// for {
	// 	select {
	// 	case <-c.ExitBuffChan:
	// 		return
	// 	}
	// }

	for range c.ExitBuffChan {
		// 得到退出消息，停止阻塞，退出Start方法
		return
	}
}

// Stop 停止/结束链接，释放资源
func (c *Connection) Stop() {
	// 1. 如果当前链接已经关闭
	if c.isClosed {
		return
	}

	c.isClosed = true

	// 关闭socket链接
	c.Conn.Close()

	// 通知从缓冲队列读数据的业务，该链接已经关闭
	c.ExitBuffChan <- true

	close(c.ExitBuffChan)
}

// GetConnID 获取当前链接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// RemoteAddr 获取远程客户端信息
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}
