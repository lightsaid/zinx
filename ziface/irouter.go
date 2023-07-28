package ziface

// IRouter 路由接口，路由包含Request信息和自定义处理业务的方法
type IRouter interface {
	PreHandle(req IRequest)  // 处理conn业务之前的钩子方法
	Handle(req IRequest)     // 处理conn业务的方法
	PostHandle(req IRequest) // 处理conn业务之后的钩子方法
}
