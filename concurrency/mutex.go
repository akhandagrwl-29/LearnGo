package main

// var (
// 	counter int
// 	wg      sync.WaitGroup
// 	mutex   sync.Mutex
// )

// func incCounter(id int) {
// 	defer wg.Done()

// 	for i := 0; i < 2; i++ {
// 		mutex.Lock()
// 		{
// 			value := counter

// 			runtime.Gosched()
// 			value++

// 			counter = value
// 		}
// 		mutex.Unlock()
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
