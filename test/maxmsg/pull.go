// package main

// import (
// 	"fmt"
// 	"time"
// 	"log"

// 	zmq "github.com/pebbe/zmq4"
// )

// func main() {
// 	// Create a ZeroMQ context
// 	context, err := zmq.NewContext()
// 	context.SetIoThreads(4)
// 	if err != nil {
// 		log.Fatal("Failed to create ZeroMQ context:", err)
// 	}

// 	// Create a PULL socket
// 	socket, err := context.NewSocket(zmq.PULL)
// 	if err != nil {
// 		log.Fatal("Failed to create PULL socket:", err)
// 	}
// 	defer socket.Close()

// 	// Connect to the PUSH server
// 	err = socket.Connect("tcp://localhost:5555")
// 	if err != nil {
// 		log.Fatal("Failed to connect to PUSH server:", err)
// 	}

// 	fmt.Println("PULL Worker Connected...")

// 	// Small delay to allow connection setup
// 	time.Sleep(time.Second)
//     i := 0
// 	for {
// 		// Receive a message (task)
// 		msg, err := socket.RecvBytes(0)
// 		if err != nil {
// 			log.Println("Error receiving message:", err)
// 			continue
// 		}

// 		fmt.Println("Received len messaage:", len(msg))
// 		fmt.Println(i)
// 		i++

// 		// Simulate task processing time
// 		time.Sleep(time.Second * 1)
// 	}
// }

package main

import (
	"fmt"
	"log"
	"os"

	zmq "github.com/pebbe/zmq4"
)

const (
	outputFilePath = "received_test.txt" // Path to save the received file
)

func main() {
	// Create a ZeroMQ context
	context, err := zmq.NewContext()
	if err != nil {
		log.Fatal("Failed to create ZeroMQ context:", err)
	}

	// Create a PULL socket
	socket, err := context.NewSocket(zmq.PULL)
	if err != nil {
		log.Fatal("Failed to create PULL socket:", err)
	}
	defer socket.Close()

	// Connect to the PUSH server
	err = socket.Connect("tcp://localhost:5555")
	if err != nil {
		log.Fatal("Failed to connect to PUSH server:", err)
	}

	fmt.Println("PULL Worker Connected...")

	// Create or truncate the output file
	file, err := os.Create(outputFilePath)
	if err != nil {
		log.Fatal("Failed to create output file:", err)
	}
	defer file.Close()

	// Receive and write chunks
	i := 0
	for {
		// Receive a chunk
		chunk, err := socket.Recv(0)
		if err != nil {
			log.Fatal("Failed to receive chunk:", err)
		}

		// Write the chunk to the file
		_, err = file.WriteString(chunk)
		if err != nil {
			log.Fatal("Failed to write chunk to file:", err)
		}

		fmt.Printf("Received chunk (%d bytes)\n", len(chunk))
		fmt.Println(i)
		i++

		// time.Sleep(time.Second * 1)
	}

	fmt.Println("File received and saved successfully as", outputFilePath)
}
