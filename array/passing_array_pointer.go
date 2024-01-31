package main

import "fmt"

func sum(array *[1e6]int) int {
	sum := 0
	for _, v := range array {
		sum += v
	}

	return sum
}

func main() {
	var array [1e6]int
	array[0] = 100

	fmt.Println(sum(&array))
}
