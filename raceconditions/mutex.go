package main

// import (
// 	"fmt"
// 	"sync"
// )

// var i int
// var m sync.Mutex

// func worker(wg *sync.WaitGroup) {
// 	m.Lock()
// 	i = i + 1
// 	m.Unlock()
// 	wg.Done()
// }

// func main() {
// 	var wg sync.WaitGroup

// 	for i := 0; i < 1000; i++ {
// 		wg.Add(1)
// 		go worker(&wg)
// 	}

// 	wg.Wait()

// 	fmt.Println("Value of i after 1000 operataions is", i)
// }
