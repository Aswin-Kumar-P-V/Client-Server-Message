package main

import (
	"encoding/json"
	"fmt"
	"net"
)

// Person type defines a person using a name and an age
type person struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

// String func implements the Stringer interface which make printing of type person convenient
func (p person) String() string {
	return fmt.Sprintf("Hello %v I see that you are %v years old,", p.Name, p.Age)
}

// HandleConnection func handles all incoming requests fron various clients
func handleConnection(conn net.Conn) {
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	pJson := string(buffer[:n])
	var p person
	json.Unmarshal([]byte(pJson), &p)

	var response string
	if p.Age > 18 {
		fmt.Println(p)
		response = fmt.Sprintln(p, "Congrats you are eligible to apply for your Driving License")
	} else {
		response = fmt.Sprintln(p, "Sorry you are not eligible to apply for a Driving license")
	}
	fmt.Println("Response To Be Send:", response)
	conn.Write([]byte(response))
	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
	fmt.Println("Server started, listening on - localhost:1234")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error Accepting Connection:", err)
			return
		}
		go handleConnection(conn)
	}
}
