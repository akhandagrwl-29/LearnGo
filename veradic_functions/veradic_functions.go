package main

import "fmt"

func getMultiples(multiplier int, args ...int) []int {
	result := make([]int, len(args))
	for i := 0; i < len(args); i++ {
		result[i] = args[i] * multiplier
		args[i]++

	}
	fmt.Println(args)
	return result
}

func main() {
	s := []int{10, 20, 30}
	a, b, c := 1, 2, 3

	mult1 := getMultiples(2, s...)
	mult2 := getMultiples(2, a, b, c)

	fmt.Println(mult1)
	fmt.Println(mult2)
	fmt.Println(s)
	fmt.Println(a, b, c)

}
