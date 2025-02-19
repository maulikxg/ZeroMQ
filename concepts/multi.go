package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func processMessage(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Simulating random processing time
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	fmt.Printf("Processed Message %d\n", id)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup
	messages := 10

	for i := 1; i <= messages; i++ {
		wg.Add(1)
		go processMessage(i, &wg) // Messages are processed concurrently
	}

	wg.Wait()
	fmt.Println("All messages processed")
}
