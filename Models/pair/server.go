package main

import (
	"fmt"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	// Create a ZeroMQ context
	context, _ := zmq.NewContext()
	socket, _ := context.NewSocket(zmq.PAIR)
	defer socket.Close()

	// Bind the PAIR socket to a TCP endpoint
	socket.Bind("tcp://*:5555")
	fmt.Println("PAIR Server Started...")

	for {
		// Send message to the client
		socket.Send("Hello from Server", 0)
		
		// Receive message from the client
		msg, _ := socket.Recv(0)
		fmt.Println("Received from client:", msg)

		time.Sleep(1 * time.Second) // Simulating processing delay
	}
}
