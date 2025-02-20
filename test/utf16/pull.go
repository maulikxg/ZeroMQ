package main

import (
	"fmt"
	"log"
	"os"

	zmq "github.com/pebbe/zmq4"
)

const (
	outputFilePath = "received_utf16_ab_4gb.txt"
)

func main() {
	context, err := zmq.NewContext()
	if err != nil {
		log.Fatal("Failed to create ZeroMQ context:", err)
	}

	socket, err := context.NewSocket(zmq.PULL)
	if err != nil {
		log.Fatal("Failed to create PULL socket:", err)
	}
	defer socket.Close()

	err = socket.Connect("tcp://localhost:5555")
	if err != nil {
		log.Fatal("Failed to connect to PUSH server:", err)
	}

	fmt.Println("PULL Worker Connected...")

	file, err := os.Create(outputFilePath)
	if err != nil {
		log.Fatal("Failed to create output file:", err)
	}
	defer file.Close()

	i := 0
	for {
		chunk, err := socket.RecvBytes(0)
		if err != nil {
			log.Fatal("Failed to receive chunk:", err)
		}

		_, err = file.Write(chunk)
		if err != nil {
			log.Fatal("Failed to write chunk to file:", err)
		}

		fmt.Printf("Received chunk %d (%d bytes)\n", i, len(chunk))
		i++
	}

	fmt.Println("File received and saved successfully as", outputFilePath)
}
