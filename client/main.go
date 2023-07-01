package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type person struct {
	Name string
	Age  int
}

func main() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error connecting to server:", err.Error())
		return
	}
	defer conn.Close()
	var pName string
	var pAge int
	fmt.Print("Enter you name:")
	fmt.Scan(&pName)
	fmt.Print("Enter your age:")
	fmt.Scan(&pAge)
	p := person{
		Name: pName,
		Age:  pAge,
	}
	pJson, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	conn.Write([]byte(string(pJson)))

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading from server:", err.Error())
		return
	}

	response := string(buffer[:n])
	fmt.Println("Server response:", response)
}
