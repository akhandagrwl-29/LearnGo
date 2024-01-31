package main

// func main() {
// 	runtime.GOMAXPROCS(runtime.NumCPU()) // different results on assigning different count of logical processors
// 	var wg sync.WaitGroup
// 	wg.Add(2)

// 	fmt.Println("start goroutines")

// 	go func() {
// 		defer wg.Done()

// 		for count := 0; count < 3; count++ {
// 			for char := 'a'; char <= 'z'; char++ {
// 				fmt.Printf("%c ", char)
// 			}
// 		}
// 	}()

// 	go func() {
// 		defer wg.Done()
// 		for count := 0; count < 3; count++ {
// 			for char := 'A'; char <= 'Z'; char++ {
// 				fmt.Printf("%c ", char)
// 			}
// 		}
// 	}()

// 	fmt.Println("Waiting to finish")
// 	wg.Wait()

// 	fmt.Println("\nTerminating program")
// }
