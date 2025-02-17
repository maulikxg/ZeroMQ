package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	numClients := 10000  // Change this to test different loads
	testDuration := 10 // Test duration in seconds

	var clients []*exec.Cmd

	fmt.Printf("Starting %d clients...\n", numClients)

	for i := 0; i < numClients; i++ {
		cmd := exec.Command("go", "run", "client.go")
		cmd.Stdout = nil // Suppress output
		cmd.Stderr = nil
		err := cmd.Start()
		if err != nil {
			fmt.Printf("❌ Failed to start client %d: %v\n", i+1, err)
			break
		}
		clients = append(clients, cmd)
		fmt.Printf("✅ Client %d started\n", i+1)
		time.Sleep(50 * time.Millisecond) // Small delay to avoid overwhelming the system
	}

	fmt.Printf("\n🕒 Test running for %d seconds...\n", testDuration)
	time.Sleep(time.Duration(testDuration) * time.Second)

	// Terminate all clients
	fmt.Println("\n🛑 Stopping all clients...")
	for i, cmd := range clients {
		if err := cmd.Process.Kill(); err != nil {
			fmt.Printf("❌ Failed to stop client %d: %v\n", i+1, err)
		} else {
			fmt.Printf("✅ Client %d stopped\n", i+1)
		}
	}

	fmt.Printf("\n🎯 Test completed: %d clients created and closed.\n", len(clients))
}
