# Simple Hello Message

<p> So basically In this example, the server listens on localhost:8080 for incoming TCP connections. When a client connects, the server reads the message sent by the client, prints it, and sends a response back to the client.

The client connects to the server using net.Dial and sends a message. It then reads the response from the server and prints it.

To run this example, you can start the server in one terminal window by running go run server.go and then start the client in another terminal window by running go run client.go. You should see the server and client exchanging messages.</p>

## Server Program:

>package main		
>
>import (		
>	"fmt"		
>	"net"		
>)		
>
>func handleConnection(conn net.Conn) {		
>	buffer := make([]byte, 1024)		
>	n, err := conn.Read(buffer)		
>	if err != nil {		
>		fmt.Println("Error reading:", err.Error())		
>		return		
>	}		
>
>	message := string(buffer[:n])		
>	fmt.Println("Received message:", message)		
>
>	response := "Hello, client!"		
>	conn.Write([]byte(response))		
>	conn.Close()		
>}		
>
>func main() {		
>	listener, err := net.Listen("tcp", "localhost:8080")		
>	if err != nil {		
>		fmt.Println("Error starting the server:", err.Error())		
>		return		
>	}		
>
>	fmt.Println("Server started. Listening on localhost:8080")		
>
>	for {		
>		conn, err := listener.Accept()		
>		if err != nil {		
>			fmt.Println("Error accepting connection:", err.Error())		
>			return		
>		}		
>
>		go handleConnection(conn)		
>	}		
>}

## Client Program

>package main
>
>import (
>	"fmt"
>	"net"
>)
>
>func main() {
>	conn, err := net.Dial("tcp", "localhost:8080")
>	if err != nil {
>		fmt.Println("Error connecting to server:", err.Error())
>		return
>	}
>	defer conn.Close()
>
>	message := "Hello, server!"
>	conn.Write([]byte(message))
>
>	buffer := make([]byte, 1024)
>	n, err := conn.Read(buffer)
>	if err != nil {
>		fmt.Println("Error reading from server:", err.Error())
>		return
>	}
>
>	response := string(buffer[:n])
>	fmt.Println("Server response:", response)
>}
