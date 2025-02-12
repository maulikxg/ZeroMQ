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

	// Connect the PAIR socket to the server
	socket.Connect("tcp://localhost:5555")
	fmt.Println("PAIR Client Connected...")

	for {
		// Receive message from the server
		msg, _ := socket.Recv(0)
		fmt.Println("Received from server:", msg)

		// Send message back to server
		socket.Send("Hello from Client", 0)

		time.Sleep(1 * time.Second)
	}
}
