package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func squareWorker(tasks <-chan int, results chan<- int, instance int, wg *sync.WaitGroup) {
// 	for num := range tasks {
// 		time.Sleep(time.Millisecond)
// 		fmt.Printf("[worker %v] Sending the result by worker %v\n", instance, instance)
// 		results <- num * num
// 	}
// 	wg.Done()
// }

// func main() {
// 	fmt.Println("[main] main started")

// 	var wg sync.WaitGroup
// 	tasks := make(chan int, 10)
// 	results := make(chan int, 10)

// 	for i := 0; i < 3; i++ {
// 		wg.Add(1)
// 		go squareWorker(tasks, results, i, &wg)
// 	}

// 	for i := 0; i < 5; i++ {
// 		tasks <- i * 2
// 	}

// 	fmt.Println("[main] main wrote 5 tasks")

// 	close(tasks)

// 	wg.Wait()

// 	for i := 0; i < 5; i++ {
// 		result := <-results
// 		fmt.Println("[main] Result", i, ":", result)
// 	}

// 	fmt.Println("[main] main stopped")
// }
