package main

import (
	"fmt"
	"runtime"
	"os"
	"time"
)


func main() {
	runtime.GOMAXPROCS(4)

	f, _ := os.Create("./log.txt")
	f.Close()

	logCh  := make(chan string , 50)

	go func() {
		for {
			msg, err := <- logCh
			if ok  {
				fmt.Println("the msg is "  + msg)
				f, _ := os.OpenFile("/tmp/log.txt", os.O_APPEND, os.ModeAppend)
				logTime := time.Now().Format(time.RFC3339)
				f.WriteString(logTime + " - " + msg)
				f.Close()
			} else {
				break
			}
		}
	}()




	mutex := make(chan bool, 1)	
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			mutex <- true
			go func() {
				//msg := fmt.Sprintf("%d + %d = %d\n", i, j, i + j)
				msg := "! + 1 = 2"
				//fmt.Println(msg)
				logCh <- msg
				//fmt.Print(msg)
				<- mutex
			}()
		}
	}

	fmt.Scanln()
}


