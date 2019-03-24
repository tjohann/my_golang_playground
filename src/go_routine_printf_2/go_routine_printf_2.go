// printf via go-routines and channels
package main

import (
	"fmt"
)

func producer(text chan string) {
	text <- "one"
	text <- "two"
	text <- "three"
	close(text)
}

func print(text chan string, done chan bool) {
	var s string
	
	for ok := true; ok; s, ok = <-text {
		fmt.Println(s)
	}
	
	done <- true
}

func main() {

	channel := make(chan string)
	done := make(chan bool)
	
	go producer(channel)
	go print(channel, done)
	
	<-done
}
