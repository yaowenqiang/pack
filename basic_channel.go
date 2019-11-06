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
	close(ch)

	for i := 0; i < len(words); i ++ {
		fmt.Print(<-ch2 + " ")
	}

	ch <- "test"
}

