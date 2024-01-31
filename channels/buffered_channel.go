package main

//func squares(c chan int) {
//	for i := 0; i < 1; i++ {
//		num := <-c
//		fmt.Println(num * num)
//	}
//	fmt.Println("I am done")
//}
//
////func greet(c chan int) {
////	fmt.Println("Hey Avantika", <-c)
////}
//
//func main() {
//	fmt.Println("Main() started")
//	c := make(chan int, 3)
//	go squares(c)
//
//	c <- 1
//	c <- 2
//	c <- 3
//	c <- 4
//	c <- 5
//
//	fmt.Println(len(c))
//
//	//fmt.Println("Hey ?")
//
//	fmt.Println("Main() stopped")
//}
