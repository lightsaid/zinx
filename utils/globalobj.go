package utils

import (
	"encoding/json"
	"os"

	"github.com/lightsaid/zinx/ziface"
)

// 定义一个全局对象
var GlobalObject *GlobalObj

// GlobalObj zinx 全局参数配置
type GlobalObj struct {
	TcpServer     ziface.IServer // 当前zinx server 全局对象
	Host          string         // 当前服务器主机IP
	TcpPort       int            // 当前服务器监听的端口
	Name          string         // 当前服务器名称
	Version       string         // 当前zinx版本
	MaxPacketSize uint32         // 传输数据包最大值
	MaxConn       int            // 当前服务器允许链接最大数
}

func init() {
	// 初始化GlobalObject，设置默认值
	GlobalObject = &GlobalObj{
		Name:          "ZinxServerApp",
		Version:       "V0.4",
		TcpPort:       8000,
		Host:          "0.0.0.0",
		MaxConn:       12000,
		MaxPacketSize: 4096,
	}

	// 从配置文件中加载一些用户配置参数

}

// Reload 加载配置文件
func (g *GlobalObj) Reload() {
	data, err := os.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}

	// 将JSON数据解析到struct中
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}
