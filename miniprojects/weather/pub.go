package main 

import (
	"fmt"
	"math/rand"
	"time"
	zmq "github.com/pebbe/zmq4"
)

func main() {
	context , _ := zmq.NewContext()
	socket, _ := context.NewSocket(zmq.PUB)
	defer context.Term()
	defer socket.Close()
	socket.Bind("tcp://*:5555")
	socket.Bind("ipc://weather.ipc")

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Loop and generate random weather updates
	for {
		// zipcode := rand.Intn(100000)
		temperature := rand.Intn(215) - 80
		realhumidity := rand.Intn(50) + 10

		msg := fmt.Sprintf("37001 %d %d", temperature, realhumidity)
		// msg := fmt.Sprintf("%05d %d %d", zipcode, temperature, realhumidity)


		socket.Send(msg, 0)
	}
	
}

