package main

// var wg sync.WaitGroup

// func Runner(baton chan int) {
// 	var newRunner int
// 	runner := <-baton

// 	fmt.Printf("Runner %d running with Baton\n", runner)

// 	if runner != 4 {
// 		newRunner = runner + 1
// 		fmt.Printf("Runner %d to the line\n", newRunner)
// 		go Runner(baton)
// 	}

// 	time.Sleep(100 * time.Millisecond)

// 	if runner == 4 {
// 		fmt.Printf("Runner %d finished, Race over\n", runner)
// 		wg.Done()
// 		return
// 	}

// 	fmt.Printf("Runner %d exchange with runner %d\n", runner, newRunner)

// 	baton <- newRunner

// }

// func main() {
// 	baton := make(chan int)

// 	wg.Add(1)

// 	go Runner(baton)

// 	baton <- 1

// 	wg.Wait()
// }
