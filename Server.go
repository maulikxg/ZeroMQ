package main

import (
	"fmt"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	// Create a ZeroMQ context
	context, _ := zmq.NewContext()
	defer context.Term()

	// Create a REP socket
	socket, _ := context.NewSocket(zmq.REP)
	defer socket.Close()

	// Bind the socket to a TCP port
	socket.Bind("tcp://*:5555")

	fmt.Println("Server is running...")

	for {
		// Wait for a request from the client
		msg, _ := socket.Recv(0)
		fmt.Printf("Received: %s\n", msg)

		// Simulate some work
		time.Sleep(time.Second)

		// Send a reply back to the client
		socket.Send("World", 0)
	}
}