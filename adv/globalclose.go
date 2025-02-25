import main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	"context"
	"log"
	zmq "github.com/pebbe/zmq4"
)


func recvMessages(ctx context.Context, socket *zmq.Socket, wg *sync.WaitGroup) {
	defer wg.Done()
	poller := zmq.NewPoller()
	poller.Add(socket, zmq.POLLIN)

	fmt.Println("Receiver started...")

	for {
		select {
		case <-ctx.Done(): // Stop if global shutdown is triggered
			fmt.Println("Receiver shutting down...")
			return
		default:
			polled, err := poller.Poll(500 * time.Millisecond) // Poll with timeout
			if err != nil {
				log.Println("Poller error:", err)
				return
			}

			for _, item := range polled {
				if item.Socket == socket {
					msg, err := socket.Recv(0)
					if err == nil {
						fmt.Println("Received:", msg)
					}
				}
			}
		}
	}
}


func main()  {
	//contex for the graceful closing
	ctx , cancel := context.WithCancel()

	var wg sync.WaitGroup

	//context and socket
	zmqcontex := zmq.NewContext()
	socket, _ := zmqcontex.NewSocket(zmq.PULL)
	defer socket.Close()
	defer zmqcontex.Term()

	socket.Bind("tcp://*:5555")

	sigChan := make(chan os.Signal, 1)
	Signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// starting receiver goroutine
	wg.Add(1)
	go receiver(ctx,socket , &wg)

	<-sigChan
	fmt.Println("Signal received, closing the socket")

	// global shutdown
	cancel()
	wg.Wait()

	fmt.Println("Graceful shutdown")

}