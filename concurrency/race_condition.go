package main

// var (
// 	counter int
// 	wg      sync.WaitGroup
// )

// func incCounter(id int) {
// 	defer wg.Done()

// 	for i := 0; i < 2; i++ {
// 		value := counter

// 		runtime.Gosched()
// 		value++

// 		counter = value
// 	}
// }

// func main() {

// 	fmt.Println("Initial value of counter:", counter)
// 	wg.Add(2)

// 	go incCounter(1)
// 	go incCounter(2)

// 	wg.Wait()
// 	fmt.Println("Final value of counter:", counter)

// }
