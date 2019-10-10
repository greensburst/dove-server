package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		fmt.Println("waiting for client message...")
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("客户端退出")
			return
		}
		fmt.Print(string(buf[:n]))
	}
}

func main() {
	fmt.Println("服务器开始监听...")
	listen, _ := net.Listen("tcp", "0.0.0.0:8888")
	defer listen.Close()

	for {
		fmt.Println("等待客户端连接...")
		conn, _ := listen.Accept()
		go process(conn)
	}
}
