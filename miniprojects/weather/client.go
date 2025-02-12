package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"os"
	"strconv"
	"strings"
)

func main() {
	context, _ := zmq.NewContext()
	socket, _ := context.NewSocket(zmq.SUB)
	defer context.Term()
	defer socket.Close()

	var temps []string
	var err error
	var temp int64
	total_temp := 0
	filter := "59937" // Default zipcode filter

	// Check for command-line argument to override the filter
	if len(os.Args) > 1 { // ./wuclient 85678
		filter = os.Args[1]
	}

	// Subscribe to the specified zipcode
	fmt.Printf("Collecting updates from weather server for %s…\n", filter)
	socket.SetSubscribe(filter)
	socket.Connect("tcp://localhost:5556")

	// Collect 101 temperature updates
	for i := 0; i < 10; i++ {
		// Receive a message from the server
		datapt, _ := socket.Recv(0)
		temps = strings.Split(datapt, " ") // Split the message into parts

		// Parse the temperature (second part of the message)
		temp, err = strconv.ParseInt(temps[1], 10, 64)
		if err == nil {
			// Add the temperature to the total
			total_temp += int(temp)
		}
	}

	// Calculate and print the average temperature
	fmt.Printf("Average temperature for zipcode %s was %dF \n\n", filter, total_temp/100)
}