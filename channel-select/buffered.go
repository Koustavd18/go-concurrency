package main

import (
	"fmt"
	"time"
)

func listenToChan(ch chan int) {

	for {
		i := <-ch
		fmt.Printf("Got %d from channel \n", i)

		time.Sleep(1 * time.Second)
	}
}

func buffered() {
	ch := make(chan int, 15)

	go listenToChan(ch)

	for i := range 100 {
		fmt.Printf("Sending %d to channel\n", i)

		ch <- i

		fmt.Println("sent")

	}
	fmt.Println("Done")
	close(ch)
}
