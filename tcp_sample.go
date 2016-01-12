// tcp server/client

package main

import (
	"fmt"
	"net"
)

func server_proc() {
	fmt.Println("Server Start...")

	listener, err := net.Listen("tcp", "localhost:40000")
	if err != nil {
		fmt.Println("Error create listener", err.Error())
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error Accept", err.Error())
			return
		}
		go doServerProc(conn)
	}
}

func doServerProc(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error Reading", err.Error())
			return
		}
		fmt.Println("Received as :%v", string(buf))
	}
}
