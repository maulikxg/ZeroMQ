package main

import (
	"fmt"
	"sync"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func connectClient(id int, wg *sync.WaitGroup, successChan chan<- int) {
	defer wg.Done()

	context, _ := zmq.NewContext()
	defer context.Term()

	// Connect to the broker
	subscriber, err := context.NewSocket(zmq.SUB)
	if err != nil {
		fmt.Printf("Client %d: Failed to create SUB socket: %v\n", id, err)
		return
	}
	defer subscriber.Close()

	err = subscriber.Connect("tcp://localhost:5556")
	if err != nil {
		fmt.Printf("Client %d: Failed to connect to broker: %v\n", id, err)
		return
	}

	subscriber.SetSubscribe("") // Subscribe to all messages

	// Signal successful connection
	successChan <- 1
	fmt.Printf("Client %d connected successfully.\n", id)

	// Keep the connection alive for a while
	time.Sleep(10 * time.Second)
}

func main() {
	var wg sync.WaitGroup
	numClients := 11000 // Number of clients to simulate
	successChan := make(chan int, numClients)

	// Simulate clients
	for i := 0; i < numClients; i++ {
		wg.Add(1)
		go connectClient(i, &wg, successChan)
	}

	// Wait for all clients to finish
	wg.Wait()

	// Count successful connections
	successCount := 0
	for i := 0; i < numClients; i++ {
		successCount += <-successChan
	}

	fmt.Printf("Total clients connected: %d\n", successCount)
}