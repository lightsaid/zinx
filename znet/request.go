package znet

import "github.com/lightsaid/zinx/ziface"

// Request 客户端请求体结构体
type Request struct {
	conn ziface.IConnection // 已经和客户端建立好的链接
	data []byte             // 客户端请求数据
}

// GetConnection 获取链接
func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

// GetData 获取客户端请求数据
func (r *Request) GetData() []byte {
	return r.data
}
