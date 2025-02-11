package main

import (
	"fmt"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	// Create a ZeroMQ context
	context, _ := zmq.NewContext()
	defer context.Term()

	// Create a REQ socket
	socket, _ := context.NewSocket(zmq.REQ)
	defer socket.Close()

	// Connect to the server
	socket.Connect("tcp://localhost:5555")

	for i := 0; i < 10; i++ {
		// Send a request
		fmt.Println("Sending Hello", i)
		socket.Send("Hello", 0)

		// Wait for a reply
		reply, _ := socket.Recv(0)
		fmt.Println("Received", reply)
	}
}