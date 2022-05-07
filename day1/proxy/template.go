package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	// 尝试启动一个server
	server, err := net.Listen("tcp", "127.0.0.1:1080")

	if err != nil {
		fmt.Println(err.Error())
		panic("服务器启动失败")
	}

	for {
		client, err := server.Accept()
		if err != nil {
			log.Printf("服务器错误：%s", err.Error())
			continue
		}
		go process(client)
	}
}

func process(conn net.Conn) {
	defer conn.Close()

	render := bufio.NewReader(conn)

	for {
		b, err := render.ReadByte()
		if err != nil {
			break
		}
		_, err = conn.Write([]byte{b})
		if err != nil {
			log.Printf("unable to write.")
			break
		}
	}
}
