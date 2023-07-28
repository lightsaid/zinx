package znet

import "github.com/lightsaid/zinx/ziface"

// 实现router时，必须实现此基础路由，
// 基础路由实现了IRouter三个接口，用户使用时，根据需要时实现具体方法即可
type BaseRouter struct{}

func (br *BaseRouter) PreHandle(req ziface.IRequest)  {}
func (br *BaseRouter) Handle(req ziface.IRequest)     {}
func (br *BaseRouter) PostHandle(req ziface.IRequest) {}

// BaseRouter 是为了提供一个已经实现 IRouter 接口结构体，当用户创建自定义Router的时候，嵌套进来即可,如：
// type MyRouter struct {
// 	BaseRouter
// }

// // Handle 仅实现业务逻辑处理方法即可，同时MyRouter也实现了IRouter接口
// func (mr *MyRouter) Handle(req ziface.HandFunc) {

// }
