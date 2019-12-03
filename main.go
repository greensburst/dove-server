package main

import (
	"dove-server/model"
	"dove-server/prc"
	"encoding/json"
	"net"
)

func main() {
	listener, _ := net.Listen("tcp", "0.0.0.0:8888") //监听8888端口
	defer listener.Close()

	for {
		conn, _ := listener.Accept() //等待连接
		go ConnectionHandler(conn)   //处理连接
	}
}

func ConnectionHandler(conn net.Conn) {
	stream := make([]byte, 1024)
	conn.Read(stream) //读取连接到stream

	var i int
	for i = 1023; i >= 0; i-- {
		if stream[i] != 0 {
			break
		}
	}

	stream = stream[:i+1]
	pkg := new(model.Package)
	json.Unmarshal(stream, pkg) //把字节流转成数据包
	prc.Processor(pkg, conn)

}
