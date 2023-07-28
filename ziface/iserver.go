package ziface

// IServer定义服务接口
type IServer interface {
	Start()                   // 启动服务
	Stop()                    // 停止服务
	Serve()                   // 开启服务方法
	AddRouter(router IRouter) // 给当前服务注册一个路由业务方法，供客户端链接处理使用
}
