package ziface

// IRequest 定义请求体接口，
// 因为将客户端的请求链接信息和请求数据都包装到 Request 中
type IRequest interface {
	GetConnection() IConnection // 获取请求体链接信息
	GetData() []byte            // 获取请求体数据
}
