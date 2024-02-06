package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4
	taskLoad         = 10
)

var wg sync.WaitGroup

func worker(tasks chan string, worker int) {
	defer wg.Done()

	for {
		task, ok := <-tasks

		if !ok {
			fmt.Printf("Worker: %d : shutting down\n", worker)
			return
		}

		fmt.Printf("Worker: %d : started %s\n", worker, task)
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		fmt.Printf("Worker: %d : completed %s\n", worker, task)

	}
}

func main() {
	tasks := make(chan string, taskLoad)
	wg.Add(numberGoroutines)

	for i := 1; i <= numberGoroutines; i++ {
		go worker(tasks, i)
	}

	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}

	close(tasks)

	wg.Wait()
}
