package main

import (
	"fmt"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	// Create a ZeroMQ context
	context, _ := zmq.NewContext()

	// Create a PULL socket
	socket, _ := context.NewSocket(zmq.PULL)
	defer socket.Close()

	// Connect to the PUSH server
	socket.Connect("tcp://localhost:5555")
	fmt.Println("PULL Worker Connected...")

	for {
		// Receive a message (task)
		msg, _ := socket.Recv(0)
		fmt.Println("Received:", msg)

		// Simulate task processing time
		time.Sleep(time.Second * 2)
	}
}
