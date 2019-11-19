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
	defer conn.Close()

	stream := make([]byte, 1024)
	conn.Read(stream) //读取连接到stream

	pkg := new(model.Package)
	json.Unmarshal(stream, pkg) //把字节流转成数据包

	proc := prc.ProcessorFactory(pkg)
	proc.Handler()

}
