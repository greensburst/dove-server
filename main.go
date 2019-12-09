package main

import (
	"dove-server/log"
	"dove-server/model"
	"dove-server/prc"
	"encoding/json"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8888") //监听8888端口
	if err != nil {
		log.Output("error.log", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept() //等待连接
		if err != nil {
			log.Output("error.log", err)
		}
		go connectionHandler(conn) //处理连接
	}
}

func connectionHandler(conn net.Conn) {
	defer conn.Close()

	stream := make([]byte, 1024)
	n, err := conn.Read(stream) //读取连接到stream
	if err != nil {
		log.Output("error.log", err)
	}
	stream = stream[:n] //删除stream多余部分

	requestPackage := new(model.RequestPackage)
	err = json.Unmarshal(stream, requestPackage) //把字节流转成数据包
	if err != nil {
		log.Output("error.log", err)
	}

	processor, err := prc.RequestPackageFactory(requestPackage) //把数据包交给包处理工厂，拿到对应的包处理器
	if err != nil {
		log.Output("error.log", err)
	}

	res, err := processor.Handler() //处理信息
	if err != nil {
		log.Output("error.log", err)
	}
	_, err = conn.Write(res) //返回响应信息给客户端
	if err != nil {
		log.Output("error.log", err)
	}
}
