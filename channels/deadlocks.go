package main

//
//func greet(c chan string) {
//	fmt.Println("Hello" + "!")
//	//time.Sleep(60 * time.Second)
//	//<-c
//	//time.Sleep(100000)
//
//}

//func main() {
//	fmt.Println("Main execution started")
//	c := make(chan string)
//
//	go greet(c)
//
//	c <- "Akhand"
//
//	//fmt.Println("Main execution stopped")
//}
//
//If you are trying to read data from a channel but channel does not have a value available with it, it blocks the current goroutine and unblocks other in a hope that some goroutine will push a value to the channel. Hence, this read operation will be blocking. Similarly, if you are to send data to a channel, it will block current goroutine and unblock others until some goroutine reads the data from it. Hence, this send operation will be blocking.
