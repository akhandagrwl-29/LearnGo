package main

//func main() {
//	fmt.Println("Main execution starts")
//
//	c := make(chan int, 3)
//
//	c <- 5
//	c <- 6
//
//	fmt.Println("The length and capacity of channel is", len(c), cap(c))
//}

//func sq(c chan int) {
//	for i := 0; i < 5; i++ {
//		fmt.Println(<-c)
//	}
//}
//
//func main() {
//	fmt.Println("Main execution starts")
//
//	c := make(chan int, 3)
//
//	go sq(c)
//
//	c <- 5
//	c <- 6
//	c <- 10
//	fmt.Println("Reached here")
//	c <- 12
//
//	fmt.Println("The length and capacity of channel is :", len(c), cap(c))
//}
