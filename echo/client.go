package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("Connection Error:", err)
	}
	defer conn.Close()

	message := "hello echo server"
	len, err := conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Write Error:", err)
	}
	fmt.Println("message length", len)

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("read Error:", err)
	}
	fmt.Println(string(buffer[:n]))
}
