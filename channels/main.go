package main

//
//import (
//	"fmt"
//	"time"
//)
//
//func read(c chan int) {
//	value := <-c
//	fmt.Println(value)
//
//	value = <-c
//	fmt.Println("Printing using below statement")
//	fmt.Println(value)
//}
//
//func main() {
//	c := make(chan int)
//
//	go read(c)
//	c <- 5
//	fmt.Println("Program reached here")
//	c <- 10
//
//	time.Sleep(100000)
//}
