package main

import (
	"fmt"
	"sync"
)

func main() {
	chPing := make(chan string)
	chPong := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for count := 0; count < 5; count++ {
			chPing <- "ping"
			m := <-chPong
			// Use channel block pringf
			fmt.Printf("Receive %s\n", m)
		}
		close(chPing)
	}()

	go func() {
		defer wg.Done()
		for m := range chPing {
			fmt.Printf("Receive %s\n", m)
			chPong <- "pong"
		}
		close(chPong)
	}()

	wg.Wait()
}
