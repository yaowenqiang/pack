package main

import (
	"time"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	godur, _  := time.ParseDuration("10ms")
		
	go func() {
		for i := 0; i < 100; i ++ {
			println("Hello")
			time.Sleep(godur)
		}
	}()
	go func() {
		for i := 0; i < 100; i ++ {
			println("Go")
			time.Sleep(godur)
		}
	}()
	dur, _  := time.ParseDuration("1s")
	time.Sleep(dur)
}



