package main

import (
	"fmt"
	"sync"
)

func SayMyName(c chan<- string, name string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		c <- name
	}
	close(c) // close from the sender side
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	var wg sync.WaitGroup
	wg.Add(3)
	go SayMyName(ch1, "ch1", &wg)
	go SayMyName(ch2, "ch2", &wg)
	go SayMyName(ch3, "ch3", &wg)

	open := 3
	for open > 0 {
		select {
		case val, ok := <-ch1:
			if !ok {
				ch1 = nil // remove from future selects
				open--
				continue
			}
			fmt.Println("message:", val)

		case val, ok := <-ch2:
			if !ok {
				ch2 = nil
				open--
				continue
			}
			fmt.Println("message:", val)

		case val, ok := <-ch3:
			if !ok {
				ch3 = nil
				open--
				continue
			}
			fmt.Println("message:", val)
		}
	}
	wg.Wait()
}
