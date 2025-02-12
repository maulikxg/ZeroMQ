package main

import (
	"fmt"
	"os"
	"bufio"

	zmq "github.com/pebbe/zmq4"
)

func main() {

	context, _ := zmq.NewContext()
	socket, _ := context.NewSocket(zmq.PUB) // publisher socket
	defer context.Term()
	defer socket.Close()

	socket.Bind("tcp://*:5555")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter for Topic: ")
		topic, _ := reader.ReadString('\n')
		fmt.Print("Enter Message: ")
		message, _ := reader.ReadString('\n')

		msg := fmt.Sprintf("%s %s", topic, message)

		socket.Send(msg, 0) 
	     
	}
	
}