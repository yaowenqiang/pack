package main

import (
	"fmt"
)

func main() {
	//messages := make(chan string, 1)
	//signals := make(chan bool, 1)

	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message: ", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "Hi"

	select {
	case messages <- msg:
		fmt.Println("send message", msg)
	default:
		fmt.Println("no message send")
	}

	select {
	case msg := <-messages:
		fmt.Println("received messages:", msg)
	case sig := <-signals:
		fmt.Println("received signal ", sig)
	default:
		fmt.Println("no activity")
	}

}
