package main

import (
	"fmt"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	// Create a ZeroMQ context
	context, _ := zmq.NewContext()

	// Create a PUSH socket
	socket, _ := context.NewSocket(zmq.PUSH)
	defer socket.Close()

	// Bind the socket to a TCP address
	socket.Bind("tcp://*:5555")
	fmt.Println("PUSH Server Started...")

	taskID := 1
	for {
		// Create a task message
		msg := fmt.Sprintf("Task #%d", taskID)
		fmt.Println("Sending:", msg)

		// Send the task to workers
		socket.Send(msg, 0)

		taskID++
		time.Sleep(time.Second) // Simulate some processing delay
	}
}
