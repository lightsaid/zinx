package main

import (
	"fmt"

	"github.com/lightsaid/zinx/ziface"
	"github.com/lightsaid/zinx/znet"
)

type MyRouter struct {
	znet.BaseRouter
}

func (mr *MyRouter) PreHandle(req ziface.IRequest) {
	fmt.Println("MyHandler ->> PreHandler")
	_, err := req.GetConnection().GetTCPConnection().Write([]byte("before handler"))
	if err != nil {
		fmt.Println("preHandle err: ", err)
	}
}

func (mr *MyRouter) Handle(req ziface.IRequest) {
	fmt.Println("MyHandler ->> Handle")
	_, err := req.GetConnection().GetTCPConnection().Write(req.GetData())
	if err != nil {
		fmt.Println("preHandle err: ", err)
	}
}

func (mr *MyRouter) PostHandle(req ziface.IRequest) {
	fmt.Println("MyHandler ->> PostHandle")
	_, err := req.GetConnection().GetTCPConnection().Write([]byte("after handler"))
	if err != nil {
		fmt.Println("PostHandle err: ", err)
	}
}

func main() {
	srv := znet.NewServer("example_zinx")
	// 注册业务处理函数
	srv.AddRouter(&MyRouter{})
	// 运行服务
	srv.Serve()

	// 等待任意输入结束
	var flag string
	fmt.Scan(&flag)
}
