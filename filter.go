package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	ch := make(chan int)
	go generate(ch)

	for {
		prime := <-ch
		fmt.Printf("%d\n", prime)
		ch1 := make(chan int)

		go filter(ch, ch1, prime)
		ch = ch1
	}
}

func generate(ch chan int) int {
	for i := 2; ; i++ {
		ch <- i
	}
}

func filter(in, out chan int, prime int) int {
	for {
		i := <-in

		if i%prime != 0 {
			out <- i

		}

	}
}
