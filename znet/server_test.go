package znet

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func ClientTest() {
	// 等待服务端启动
	time.Sleep(2 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("dial err: ", err)
		return
	}

	for {
		_, err := conn.Write([]byte("Hello Zinx"))
		if err != nil {
			fmt.Println("write to failed: ", err)
			return
		}

		buf := make([]byte, 1024)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("write to read: ", err)
			return
		}
		fmt.Printf("server call back: %s - %d \n", string(buf[:cnt]), time.Now().Second())
		time.Sleep(1 * time.Second)
	}
}

func TestServer(t *testing.T) {
	srv := NewServer("myzinx")
	srv.Serve()

	ClientTest()
}
