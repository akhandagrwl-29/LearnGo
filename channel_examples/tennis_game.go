package main

// var wg sync.WaitGroup

// func init() {
// 	rand.Seed(time.Now().UnixNano())
// }

// func player(name string, court chan int) {
// 	defer wg.Done()

// 	for {
// 		ball, ok := <-court
// 		if !ok {
// 			fmt.Printf("Player %s won\n", name)
// 			return
// 		}

// 		n := rand.Intn(100)
// 		if n%13 == 0 {
// 			fmt.Printf("Player %s missed\n", name)
// 			close(court)
// 			return
// 		}

// 		fmt.Printf("Player %s hit %d\n", name, ball)
// 		ball++

// 		court <- ball
// 	}
// }

// func main() {
// 	court := make(chan int)

// 	wg.Add(2)

// 	go player("Djokovic", court)
// 	go player("Nadal", court)

// 	court <- 1

// 	wg.Wait()

// }
