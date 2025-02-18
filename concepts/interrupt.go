package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	// Create a ZeroMQ context
	context, _ := zmq.NewContext()
	defer context.Term()

	// Create a ZeroMQ socket
	socket, _ := context.NewSocket(zmq.REP)
	defer socket.Close()

	// Bind socket to an endpoint
	socket.Bind("tcp://*:5555")

	// Create a channel to listen for OS signals
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Goroutine to handle incoming messages
	go func() {
		for {
			msg, err := socket.Recv(0)
			if err != nil {
				fmt.Println("Socket closed:", err)
				return
			}
			fmt.Println("Received:", msg)
			socket.Send("Ack", 0)
		}
	}()

	// Wait for an interrupt signal
	sig := <-sigs
	fmt.Println("Received signal:", sig)

	// Graceful cleanup
	fmt.Println("Closing ZeroMQ socket...")
	socket.Close()
	context.Term()
	fmt.Println("Shutdown complete.")
}
