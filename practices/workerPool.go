// REF: Compare with https://gobyexample.com/worker-pools
package practices

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

func worker(name string, queue chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range queue {
		fmt.Printf("Worker %s processing value: %s\n", name, val)
		// Simulate working time
		time.Sleep(5 * time.Second)
	}
}

func WorkerPool(concurrency int) {
	fmt.Println("Waiting for input..., press Ctrl+D to exit")
	queue := make(chan string)
	var wg sync.WaitGroup

	// Start worker pool
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go worker(fmt.Sprintf("worker-%d", i), queue, &wg)
	}

	// Read input from stdin and send to workers
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		queue <- scanner.Text()
	}
	close(queue) // tell workers: no more jobs

	wg.Wait()
	fmt.Println("All jobs done.")
}
