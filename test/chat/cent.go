// cent.go
package main

import (
    "fmt"
    "strings"
    "sync"
    zmq "github.com/pebbe/zmq4"
)

type Broker struct {
    usernames map[string]bool
    mutex     sync.RWMutex
}

func NewBroker() *Broker {
    return &Broker{
        usernames: make(map[string]bool),
    }
}

func (b *Broker) checkUsername(username string) bool {
    b.mutex.Lock()
    defer b.mutex.Unlock()
    
    if _, exists := b.usernames[username]; exists {
        return false
    }
    
    b.usernames[username] = true
    return true
}

func (b *Broker) removeUsername(username string) {
    b.mutex.Lock()
    defer b.mutex.Unlock()
    delete(b.usernames, username)
}

func main() {
    context, _ := zmq.NewContext()
    defer context.Term()

    broker := NewBroker()

    // Socket for publishing messages
    publisher, _ := context.NewSocket(zmq.PUB)
    defer publisher.Close()
    publisher.Bind("tcp://*:5556")

    // Socket for receiving messages
    subscriber, _ := context.NewSocket(zmq.SUB)
    defer subscriber.Close()
    subscriber.Bind("tcp://*:5555")
    subscriber.SetSubscribe("")

    fmt.Println("Central broker running...")

    for {
        // Receive message
        message, err := subscriber.Recv(0)
        if err != nil {
            continue
        }

        parts := strings.SplitN(message, ":", 3)
        if len(parts) < 2 {
            continue
        }

        command := parts[0]
        username := parts[1]

        switch command {
        case "REGISTER":
            if broker.checkUsername(username) {
                // Username is available
                publisher.Send(fmt.Sprintf("REGISTER_OK:%s:success", username), 0)
            } else {
                // Username is taken
                publisher.Send(fmt.Sprintf("REGISTER_FAIL:%s:taken", username), 0)
            }

        case "UNREGISTER":
            broker.removeUsername(username)
            publisher.Send(fmt.Sprintf("SYSTEM:all:%s has left the chat", username), 0)

        default:
            // Regular chat message, forward it
            if len(parts) == 3 {
                publisher.Send(message, 0)
            }
        }
    }
}