package main

// func printPrime(prefix string) {
// 	defer wg.Done()
// next:
// 	for outer := 2; outer < 5000; outer++ {
// 		for inner := 2; inner < outer; inner++ {
// 			if outer%inner == 0 {
// 				continue next
// 			}
// 		}
// 		fmt.Printf("%s:%d\n", prefix, outer)
// 	}
// 	fmt.Println("Completed", prefix)
// }

// var wg sync.WaitGroup

// func main() {
// 	runtime.GOMAXPROCS(1)
// 	wg.Add(2)

// 	fmt.Println("Create Goroutines")
// 	go printPrime("A")
// 	go printPrime("B")

// 	fmt.Sprintln("Waiting to finish")
// 	wg.Wait()

// 	fmt.Println("\nTerminating program")
// }
