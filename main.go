package main

import (
	"encoding/json"
	"net"
)

func main() {
	listener, _ := net.Listen("tcp", "0.0.0.0:8888")
	defer listener.Close()
	for {
		conn, _ := listener.Accept()
		go process(conn)
	}
}

func handle_connection(conn net.Conn) {

	defer conn.Close()
	for {
		receiver := make([]byte, 1024)
		conn.Read(receiver)
		datapkg := new(data_package)
		json.Unmarshal(receiver, datapkg)

		processor := new(processor)
		switch datapkg.data_header.message_code {
		case SIGNUP_INFORMATION:
			processor.signup(datapkg.content)
		}
	}
}
