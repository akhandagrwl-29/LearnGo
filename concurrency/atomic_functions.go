package main

// var (
// 	wg      sync.WaitGroup
// 	counter int64
// )

// func main() {
// 	wg.Add(2)
// 	fmt.Println("The initial value of counter is:", counter)

// 	go incCounter(1)
// 	go incCounter(2)

// 	wg.Wait()
// 	fmt.Println("The final value of counter is:", counter)
// }

// func incCounter(id int) {
// 	defer wg.Done()

// 	for i := 0; i < 2; i++ {
// 		atomic.AddInt64(&counter, 1)

// 		runtime.Gosched()
// 	}
// }
