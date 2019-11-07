package main
import (
	"fmt"
	"strings"
)

func main() {
	phrase := "These are the times that try mans's souls.\n"
	words := strings.Split(phrase, " ")

	ch := make(chan string, 1)
	ch <- "Hello"
	fmt.Println(<-ch)

	ch2 := make(chan string, len(words))
	for _, word := range words {
		ch2 <- word
	}
	close(ch2)

	/*
	for {
		if msg, ok := <- ch2; ok {
			fmt.Print(msg + " ")
		} else {
			break
		}
	}
	*/

	for msg := range ch2 {
		fmt.Print(msg + " ")
	}

	msgCh := make(chan Message, 1)
	errCh := make(chan FailedMessage, 1)


	msg := Message {
		To: []string{"jacky@test.com"},
		From: "john@test.com",
		Content: "Keep it secrect, Keep it safe.",
	}


	failedMessage := FailedMessage {
		ErrorMessage: "Message intercepted by black rider",
		OriginalMessage: Message{},
	}

	msgCh <- msg
	errCh <- failedMessage

	/*
	fmt.Println(<-msgCh)
	fmt.Println(<-errCh)
	*/

	//msgCh <- msg
//	errCh <- failedMessage

	select {
		case receivedMsg := <- msgCh:
			fmt.Println(receivedMsg)
		case receivedError := <- errCh:
			fmt.Println(receivedError)
		default:
			fmt.Println("No messages received")
	}

}


type Message struct {
	To []string
	From string
	Content string
}

type FailedMessage  struct {
	ErrorMessage string
	OriginalMessage Message
}
