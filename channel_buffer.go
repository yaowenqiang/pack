package main

import "fmt"

func main() {
	fmt.Println("before chan")
	message := make(chan string)
	//message <- "ping"
	//message <- "pang"
	//message <- "channel"
	go func() {
		//fmt.Println(<-message)
		//message <- "buffered"
		//<-message
		fmt.Println("put message")
		fmt.Println(<-message)
	}()
	go func() {
		//fmt.Println(<-message)
		fmt.Println("get message")
		message <- "buffered"
		//<-message
	}()
	fmt.Println("after chan")

	/*

		for i := 0; i < 3; i++ {

		}
	*/

	//fmt.Println(<-message)
	//fmt.Println(<-message)
	//fmt.Println(<-message)
	//fmt.Scanln()
}
