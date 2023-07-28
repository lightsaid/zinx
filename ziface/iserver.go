package ziface

// IServer定义服务接口
type IServer interface {
	Start() // 启动服务
	Stop()  // 停止服务
	Serve() //开启服务方法
}
