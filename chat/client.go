// package main 

// import (
// 	"fmt"
// 	"bufio"
// 	"strings"
// 	"os"
// 	zmq "github.com/pebbe/zmq4"
// )

// func main() {

// 	context, _ := zmq.NewContext()
// 	defer context.Term()

//     // sockets as the publisher for the send to cental server.
// 	publisher, _ := context.NewSocket(zmq.PUB)
// 	defer publisher.Close()
// 	publisher.Connect("tcp://localhost:5555")

// 	// socket as a subscriber for the listen to central server
// 	subscriber, _ := context.NewSocket(zmq.SUB)
// 	defer subscriber.Close()
// 	subscriber.Connect("tcp://localhost:5556")
// 	subscriber.SetSubscribe("") // Subscribe to all messages

// 	// username 
// 	fmt.Print("Enter your name: ")
// 	reader := bufio.NewReader(os.Stdin)
// 	username, _ := reader.ReadString('\n')
// 	username = strings.TrimSpace(username)


// 	fmt.Println("\nWelcome to the chat, " + username + "! Type 'quit' to exit.\n")

// 	// Start listening for messages in a separate goroutine
// 	go func() {
// 		for {
// 			msg, _ := subscriber.Recv(0)
// 			parts := strings.SplitN(msg, ":", 2)
// 			if len(parts) == 2 && parts[0] != username {
// 				fmt.Printf("\n%s: %s\n", parts[0], parts[1])
// 				fmt.Print("Enter message: ") // Keep input prompt consistent
// 			}
			
// 		}
// 	}()

// 	// for the send the messages 
//     // Read user input and send messages

// 	for {
//         fmt.Print("Enter message: ")
// 		message, _ := reader.ReadString('\n')
// 		message = strings.TrimSpace(message)

// 		if message == "quit" {
// 			break
// 		}

// 		publisher.Send(fmt.Sprintf("%s: %s", username, message), 0)
// 		fmt.Printf("You: %s\n", message)
// 	}
	
// 	fmt.Println("Chat ended. Goodbye!")
// }


package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	context, _ := zmq.NewContext()
	defer context.Term()

	// Create a PUB socket to send messages to the central server
	publisher, _ := context.NewSocket(zmq.PUB)
	defer publisher.Close()
	publisher.Connect("tcp://localhost:5555")

	// Create a SUB socket to receive messages from the central server
	subscriber, _ := context.NewSocket(zmq.SUB)
	defer subscriber.Close()
	subscriber.Connect("tcp://localhost:5556")
	subscriber.SetSubscribe("") // Subscribe to all messages

	// Get the username from the user
	fmt.Print("Enter your name: ")
	reader := bufio.NewReader(os.Stdin)
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Println("\nWelcome to the chat, " + username + "! Type '@username message' to send a private message.\n")

	// Start listening for messages in a separate goroutine
	go func() {
		for {
			msg, _ := subscriber.Recv(0)
			parts := strings.SplitN(msg, ":", 3) // Format: sender:targetUser:message

			if len(parts) == 3 {
				sender := strings.TrimSpace(parts[0])
				targetUser := strings.TrimSpace(parts[1])
				message := strings.TrimSpace(parts[2])

				// Show message if it's a group message or meant for this user
				if targetUser == "all" || targetUser == username {
					fmt.Printf("\n%s: %s\n", sender, message)
					fmt.Print("Enter message: ") // Keep input prompt consistent
				}
			}
		}
	}()

	// Read user input and send messages
	for {
		fmt.Print("Enter message: ")
		message, _ := reader.ReadString('\n')
		message = strings.TrimSpace(message)

		if message == "quit" {
			break
		}

		// Check if message is private (@username message)
		targetUser := "all" // Default is to send to everyone
		if strings.HasPrefix(message, "@") {
			words := strings.SplitN(message, " ", 2)
			if len(words) == 2 {
				targetUser = strings.TrimPrefix(words[0], "@") // Extract the target username
				message = words[1] // Get the actual message
			}
		}

		// Send the formatted message
		publisher.Send(fmt.Sprintf("%s:%s:%s", username, targetUser, message), 0)
		fmt.Printf("You to %s: %s\n", targetUser, message)
	}

	fmt.Println("Chat ended. Goodbye!")
}
