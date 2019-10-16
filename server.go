package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type User struct {
	Nickname string `json:"nickname"`
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

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
		if string(buf[:n]) == "islogin" {
			res := "yes"
			buf := []byte(res)
			n, err = conn.Write(buf)
			if err != nil {
				fmt.Println("服务器错误")
			}
		} else {
			user := new(User)
			json.Unmarshal(buf[:n], user)
			if user.Mail == "116@qq.com" && user.Password == "123" {
				res := "access"
				buf := []byte(res)
				n, err = conn.Write(buf)
				if err != nil {
					fmt.Println("服务器错误")
				}
			} else {
				res := "error"
				buf := []byte(res)
				n, err = conn.Write(buf)
				if err != nil {
					fmt.Println("服务器错误")
				}
			}
		}
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
