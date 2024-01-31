package main

//func greet(c chan string) {
//	fmt.Println("Hello", <-c, "!")
//}
//
//func greeter(cc chan chan string) {
//
//	fmt.Println("greeter started")
//
//	c := make(chan string)
//	cc <- c
//
//	fmt.Println("greeter stopped")
//}
//
//func main() {
//	fmt.Println("Main execution started")
//
//	cc := make(chan chan string)
//
//	go greeter(cc)
//
//	fmt.Println("Reached here")
//
//	c := <-cc
//
//	go greet(c)
//
//	c <- "Akhand"
//
//	fmt.Println("main stopped")
//}
