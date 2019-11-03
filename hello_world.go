package main

import (
)

func main() {
	go func() {
		println("Hello")
	}()
	go func() {
		println("Go")
	}()
}



