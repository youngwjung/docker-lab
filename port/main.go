package main

import (
	"net"
	"strings"
)

func main() {
	server, _ := net.Listen("tcp", ":0")
	defer server.Close()

	for {
		conn, _ := server.Accept()
		defer conn.Close()

		conn.Write([]byte("암호를 입력하세요.\n"))
		go handler(conn)
	}
}

func handler(conn net.Conn) {
	recv := make([]byte, 4096)

	for {
		n, _ := conn.Read(recv)

		message := strings.TrimSpace(string(recv[:n]))

		if message == "love dive" {
			conn.Write([]byte("암호키는 2aXQzHGUCj 입니다.\n"))
			conn.Close()
		} else {
			conn.Write([]byte("암호가 틀렸습니다.\n"))
		}
	}
}
