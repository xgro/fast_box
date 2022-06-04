package main

import (
	"fmt"
	"time"
)

func readword(ch chan string) {
	fmt.Println("Type a word, then hit Enter.")
	var word string
	fmt.Scanf("%s", &word)
	ch <- word
}

func printchar() {
	for {
		fmt.Printf(".")
		time.Sleep(2 * time.Second)
	}
}

func main() {
	defer fmt.Println("===== BYE..")
	go printchar()

	ch := make(chan string)
	go readword(ch)

	select {
	case word := <-ch:
		fmt.Println("\nReceived: ", word)
	}
}
