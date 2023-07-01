package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err.Error())
		return
	}
	defer conn.Close()
	var name string
	fmt.Println("Enter you name:")
	fmt.Scan(&name)
	conn.Write([]byte(name))

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading from server:", err.Error())
		return
	}

	response := string(buffer[:n])
	fmt.Println("Server response:", response)
}
