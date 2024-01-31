package main

import "fmt"

func greet(msg string) {
	fmt.Println(msg)
}

func main() {
	fmt.Println("Print one")
	defer greet("One")

	fmt.Println("Print Two")
	defer greet("Two")

	fmt.Println("Print Three")
	defer greet("Three")
}
