package main

import (
	"fmt"
	"log"
"os"
	zmq "github.com/pebbe/zmq4"
)

const (
	chunk = 2*1024*1024 * 1024 // 1 MB chunks (adjust as needed)
	fileSize  = 10 * 1024 * 1024 * 1024 // 4 GB file size
	filename  = "test.txt" // Path to the file to send
)

func main() {
	// Create a ZeroMQ context
	context, err := zmq.NewContext()
	if err != nil {
		log.Fatal("Failed to create ZeroMQ context:", err)
	}

	// Create a PUSH socket
	socket, err := context.NewSocket(zmq.PUSH)
	if err != nil {
		log.Fatal("Failed to create PUSH socket:", err)
	}
	defer socket.Close()

	// Bind the socket to a TCP address
	err = socket.Bind("tcp://*:5555")
	if err != nil {
		log.Fatal("Failed to bind PUSH socket:", err)
	}

	fmt.Println("PUSH Server Started...")


	// for file stuff 

	// create the new dummy file 
	fmt.Println("Creating the dummy file")
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Failed to create file:", err)
	}
	defer file.Close()


	// write the data to the file
	fmt.Println("Writing the data to the file")
	dummyData := make([]byte, chunk)

	for i := 0; i < fileSize/chunk; i++ {
		_, err = file.Write(dummyData)
		if err != nil {
			log.Fatal("Failed to write data to file:", err)
		}
	}

	fmt.Println("File written successfully")

	// open file for reading 
	file , err = os.Open(filename)
	if err != nil {
		log.Println("Error opening file")
	}

	// read the data from the file
	buffer := make([]byte, chunk)
	chunkID := 0

	for {
		n , err := file.Read(buffer)
		if err != nil {
			if err.Error() == "EOF" {
				break // End of file reached
			}
			log.Fatal("Failed to read file:", err)
		}

		// Send the chunk
		_, err = socket.SendBytes(buffer[:n], 0)
		if err != nil {
				log.Fatal("Failed to send chunk:", err)
		}
		
		fmt.Printf("Sent chunk %d (%d bytes)\n", chunkID, n)
		chunkID++
	}

	fmt.Println("File sent successfully")


	// n := 10
	// for i:=0; i<n; i++ {
    //     // sending the message no.
	// 	fmt.Println("Sending no:", i)

	// 	// Send a message (task)
	// 	msg := fmt.Sprintf("Message no:%d and byte message : %s", i , make([]byte, 2*1024*1024*1024))

	// 	_, err = socket.Send(msg, 0)
	// 	if err != nil {
	// 		log.Println("Error sending message:", err)
	// 	}

	// }

	// for file stuff 


}
