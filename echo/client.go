package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9090")
	if err != nil {
		fmt.Println("Connection Error:", err)
		return
	}
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		if err != io.EOF {
			fmt.Println("Reading Error:", err)
			return
		}
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("current time:", string(buffer[:n]))
}
