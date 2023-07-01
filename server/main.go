package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	name := string(buffer[:n])
	fmt.Println("Hello", name)

	response := fmt.Sprintln("Hello ", name)
	conn.Write([]byte(response))
	conn.Close()

}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Oops Error Starting the server:", err)
		return
	}
	fmt.Println("Server started, listening on - localhost:8080")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting the connection:", err)
			return
		}
		go handleConnection(conn)

	}

}
