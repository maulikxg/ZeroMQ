package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"time"

	zmq "github.com/pebbe/zmq4"
)

const (
	chunkSize = 2 * 1024 * 1024        // 2MB chunks
	fileSize  = 4 * 1024 * 1024 * 1024 // 4GB file
	filename  = "utf16_ab_4gb.txt"
)

func main() {
	context, err := zmq.NewContext()
	if err != nil {
		log.Fatal("Failed to create ZeroMQ context:", err)
	}

	socket, err := context.NewSocket(zmq.PUSH)
	if err != nil {
		log.Fatal("Failed to create PUSH socket:", err)
	}
	defer socket.Close()

	err = socket.Bind("tcp://*:5555")
	if err != nil {
		log.Fatal("Failed to bind PUSH socket:", err)
	}

	fmt.Println("PUSH Server Started...")

	// Create UTF-16 file
	fmt.Println("Creating a 4GB UTF-16 file filled with 'ab'...")
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Failed to create file:", err)
	}
	defer file.Close()

	// Write UTF-16 BOM (Little Endian)
	file.Write([]byte{0xFF, 0xFE})

	// Create a buffer filled with "ab" in UTF-16 (Little Endian)
	buffer := make([]byte, chunkSize)
	for i := 0; i < chunkSize/4; i++ {
		binary.LittleEndian.PutUint16(buffer[i*4:], 'a')   // 'a' -> 0x0061
		binary.LittleEndian.PutUint16(buffer[i*4+2:], 'b') // 'b' -> 0x0062
	}

	// Write buffer repeatedly to reach 4GB
	for i := 0; i < fileSize/chunkSize; i++ {
		_, err = file.Write(buffer)
		if err != nil {
			log.Fatal("Failed to write data to file:", err)
		}
	}

	fmt.Println("UTF-16 file created successfully.")

	// Open the file for reading
	file, err = os.Open(filename)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	// Send file in chunks
	readBuffer := make([]byte, chunkSize)
	chunkID := 0

	for {
		n, err := file.Read(readBuffer)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatal("Failed to read file:", err)
		}

		_, err = socket.SendBytes(readBuffer[:n], 0)
		if err != nil {
			log.Fatal("Failed to send chunk:", err)
		}

		fmt.Printf("Sent chunk %d (%d bytes)\n", chunkID, n)
		chunkID++

		time.Sleep(time.Millisecond * 500) // Debugging delay
	}

	fmt.Println("File sent successfully.")
}
