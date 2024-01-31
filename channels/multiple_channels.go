package main

//func square(c chan int) {
//	fmt.Println("square reading")
//	num := <-c
//	fmt.Println("square reading done")
//	c <- num * num
//}
//
//func cube(c chan int) {
//	fmt.Println("cube reading")
//	num := <-c
//	fmt.Println("cube reading done")
//	c <- num * num * num
//
//}
//
//func main() {
//	fmt.Println("Main execution starts")
//
//	squareChan := make(chan int)
//	cubeChan := make(chan int)
//
//	go square(squareChan)
//	go cube(cubeChan)
//
//	time.Sleep(5 * time.Second)
//
//	testNum := 3
//	fmt.Println("main sent testNum to squareChan")
//
//	squareChan <- testNum
//
//	fmt.Println("main resuming")
//	fmt.Println("main sent testNum to cubeChan")
//
//	cubeChan <- testNum
//
//	fmt.Println("main resuming")
//	fmt.Println("main reading from channels")
//
//	squareVal := <-cubeChan
//	cubeVal := <-squareChan
//	sum := squareVal + cubeVal
//
//	fmt.Println("The sum of square and cube of", testNum, "is", sum)
//	fmt.Println("main execution finished")
//}
