package main

import (
	"fmt"
	"strings"
)

func shout(ping <-chan string, pong chan<- string) {
	for {

		s := <-ping

		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))

	}
}

func PingPong() {

	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	fmt.Println("Type something and press enter (q to quit)....")

	for {
		fmt.Print("->")

		var userInput string

		fmt.Scanln(&userInput)

		if userInput == strings.ToLower("q") {
			break
		}

		ping <- userInput

		response := <-pong
		fmt.Println(response)
	}

	fmt.Println("Closing channels")
	close(ping)
	close(pong)
}
