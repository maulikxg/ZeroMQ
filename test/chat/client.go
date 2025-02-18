// client.go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "time"
    zmq "github.com/pebbe/zmq4"
)

func main() {
    context, _ := zmq.NewContext()
    defer context.Term()

    var username string

    // Socket to send messages
    publisher, _ := context.NewSocket(zmq.PUB)
    defer publisher.Close()
    publisher.Connect("tcp://localhost:5555")

    // Socket to receive messages
    subscriber, _ := context.NewSocket(zmq.SUB)
    defer subscriber.Close()
    subscriber.Connect("tcp://localhost:5556")
    subscriber.SetSubscribe("")

    // Wait for connection to establish
    time.Sleep(time.Second)

    reader := bufio.NewReader(os.Stdin)
    
    // Username registration
    for {
        fmt.Print("Enter your name: ")
        username, _ = reader.ReadString('\n')
        username = strings.TrimSpace(username)

        if username == "" {
            fmt.Println("Username cannot be empty")
            continue
        }

        // Send registration request
        publisher.Send(fmt.Sprintf("REGISTER:%s:request", username), 0)

        // Wait for response
        for {
            message, err := subscriber.Recv(0)
            if err != nil {
                continue
            }

            parts := strings.SplitN(message, ":", 3)
            if len(parts) != 3 {
                continue
            }

            command := parts[0]
            user := parts[1]

            if user != username {
                continue
            }

            if command == "REGISTER_OK" {
                goto CHAT_START
            } else if command == "REGISTER_FAIL" {
                fmt.Println("Username already taken. Please choose another one.")
                break
            }
        }
    }

CHAT_START:
    fmt.Printf("\nWelcome to the chat, %s!\nType '@username message' for private messages or 'quit' to exit.\n\n", username)

    // Start message receiver
    go func() {
        for {
            message, err := subscriber.Recv(0)
            if err != nil {
                continue
            }

            parts := strings.SplitN(message, ":", 3)
            if len(parts) != 3 {
                continue
            }

            sender := parts[0]
            target := parts[1]
            content := parts[2]

            // Skip own messages
            if sender == username {
                continue
            }

            // Show message if it's for everyone or specifically for this user
            if target == "all" || target == username {
                if sender == "SYSTEM" {
                    fmt.Printf("\n[System] %s\n", content)
                } else {
                    fmt.Printf("\n%s: %s\n", sender, content)
                }
                fmt.Print("Enter message: ")
            }
        }
    }()

    // Message sending loop
    for {
        fmt.Print("Enter message: ")
        message, _ := reader.ReadString('\n')
        message = strings.TrimSpace(message)

        if message == "quit" {
            publisher.Send(fmt.Sprintf("UNREGISTER:%s:leaving", username), 0)
            break
        }

        if message == "" {
            continue
        }

        targetUser := "all"
        if strings.HasPrefix(message, "@") {
            parts := strings.SplitN(message, " ", 2)
            if len(parts) == 2 {
                targetUser = strings.TrimPrefix(parts[0], "@")
                message = parts[1]
            }
        }

        publisher.Send(fmt.Sprintf("%s:%s:%s", username, targetUser, message), 0)
    }

    fmt.Println("Chat ended. Goodbye!")
}