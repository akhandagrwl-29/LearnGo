package main

//func sender(c chan int) {
//	c <- 1
//	c <- 2
//	c <- 3
//	c <- 4
//	close(c)
//}
//
//func main() {
//	c := make(chan int, 3)
//
//	go sender(c)
//
//	fmt.Printf("Length of channle is %v and capacity of channel is %v\n", len(c), cap(c))
//
//	for val := range c {
//		fmt.Printf("Length of channel after value %v read is %v\n", val, len(c))
//		_, ok := <-c
//
//		fmt.Println(ok)
//	}
//
//	//_, ok := <-c
//	//
//	//fmt.Println(ok)
//}
