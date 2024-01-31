package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

func service1(c chan string) {
	fmt.Println("service1 started", time.Since(start))
	//time.Sleep(3 * time.Second)
	time.Sleep(2 * time.Second)
	c <- "Hello from service 1"

}

func service2(c chan string) {
	fmt.Println("service2 started", time.Since(start))
	//time.Sleep(5 * time.Second)
	c <- "Hello from service 2"
}

func main() {
	fmt.Println("Main execution started")

	chan1 := make(chan string)
	chan2 := make(chan string)

	go service1(chan1)
	go service2(chan2)

	fmt.Println("Program reached here")

	time.Sleep(1 * time.Second)

	select {
	case res := <-chan2:
		fmt.Println("Response from service 2", res, time.Since(start))
	case res := <-chan1:
		fmt.Println("Response from service 1", res, time.Since(start))
	default:
		fmt.Println("we get some anonymous response", time.Since(start))
	}

	fmt.Println("Main execution stopped", time.Since(start))
}
