package main

import "fmt"

func main() {
	message := make(chan string)
	//message <- "buffered"
	//message <- "channel"

	go func() {
		message <- "ping"
		message <- "pang"
	}()

	fmt.Println(<-message)
	fmt.Println(<-message)
	//fmt.Println(<-message)
	//fmt.Scanln()
}
