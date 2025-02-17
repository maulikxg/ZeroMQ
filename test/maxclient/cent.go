package main

import (
	"fmt"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	// Create ZeroMQ context
	context, _ := zmq.NewContext()
	defer context.Term()

	// Create an XSUB socket (receives messages from clients)
	xsub, _ := context.NewSocket(zmq.XSUB)
	defer xsub.Close()
	xsub.Bind("tcp://*:5555") // Clients send messages here

	// Create an XPUB socket (sends messages to clients)
	xpub, _ := context.NewSocket(zmq.XPUB)
	defer xpub.Close()
	xpub.Bind("tcp://*:5556") // Clients receive messages from here

	fmt.Println("Central broker running...")

	// Forward messages between XSUB and XPUB
	zmq.Proxy(xsub, xpub, nil)
}
