package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		url := strings.ToUpper(scanner.Text())
		if url == "q" {
			fmt.Println("Quit")
			return
		} else {
			fmt.Println(url)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "err:", err)
		os.Exit(1)

	}
}
