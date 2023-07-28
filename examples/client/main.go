package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	fmt.Println("client test")

	// 等待服务端启动
	time.Sleep(2 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal("dial to failed: ", err)
	}

	for {
		// 写入
		_, err := conn.Write([]byte("Zinx v0.3"))
		if err != nil {
			fmt.Println("write error: ", err)
			return
		}

		// 读取
		buf := make([]byte, 1024)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read failed: ", err)
			return
		}

		fmt.Printf("read data: %s \n", string(buf[:cnt]))
		time.Sleep(time.Second)
	}
}
