package main

import (
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

	for {
		stream := make([]byte, 1024)
		conn.Read(stream) //读取连接到stream

		pkg := new(Package)
		json.Unmarshal(stream, pkg) //把字节流转成数据包

		switch pkg.Header.Code { //判断消息码
		case SIGNUP_MESSAGE: //如果是注册消息
			up := new(prc.UserPrc) //实例化一个user处理器
			up.Signup(pkg.Body)    //把包体信息传给注册方法
		case SIGNIN_MESSAGE:

		}
	}
}
