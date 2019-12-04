package main

import (
	"dove-server/model"
	"dove-server/prc"
	"encoding/json"
	"fmt"
	"net"
)

func main() {
	listener, _ := net.Listen("tcp", "0.0.0.0:8888") //监听8888端口
	defer listener.Close()

	for {
		conn, _ := listener.Accept() //等待连接
		go connectionHandler(conn)   //处理连接
	}
}

func connectionHandler(conn net.Conn) {
	defer conn.Close()

	stream := make([]byte, 1024)
	conn.Read(stream) //读取连接到stream

	var i int
	for i = 1023; i >= 0; i-- {
		if stream[i] != 0 {
			break
		}
	}
	stream = stream[:i+1] //删除stream多余部分

	requestPackage := new(model.RequestPackage)
	json.Unmarshal(stream, requestPackage) //把字节流转成数据包

	processor := prc.RequestPackageFactory(requestPackage) //把数据包交给包处理工厂，拿到对应的包处理器
	res, err := processor.Handler()                        //处理信息
	if err != nil {
		fmt.Println(err)
	}
	conn.Write(res)
}
