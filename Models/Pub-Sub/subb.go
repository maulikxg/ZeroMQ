package main 

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"bufio"
	"os"
	"strings"
)

func main() {
	context, _ := zmq.NewContext()
	socket, _ := context.NewSocket(zmq.SUB)
	defer context.Term()
	defer socket.Close()
	socket.Connect("tcp://localhost:5555")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Topic u want to subscribe: ")
	topic, _ := reader.ReadString('\n')
	
	socket.SetSubscribe(topic)

	for {
		msg, _ := socket.Recv(0)
				// Split the message to separate topic and actual message
				parts := strings.SplitN(msg, " ", 2)
				if len(parts) < 2 {
					fmt.Println("Invalid message format:", msg)
					continue
				}
		
				// Extract the message (second part)
				message := parts[1]
		
				// Print only the actual message
				fmt.Println("Received Message:", message)
		
		
	}
	

}