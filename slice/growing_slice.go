package main

// func main() {

//  Example-1: When capcacity of new slice have space, so original
// underlying array gets modified on append
// slice := []int{10, 20, 30, 40, 50}
// newSlice := slice[1:3]
// fmt.Println(slice)
// fmt.Println(newSlice)
// newSlice = append(newSlice, 45)
// fmt.Println(slice)
// fmt.Println(newSlice)

//  Example-2: When capcacity of new slice have no space, append will change the underlying array of new slice.
// slice := []int{10, 20, 30, 40, 50}
// newSlice := slice[1:5]
// fmt.Println(slice)
// fmt.Println(newSlice)
// newSlice = append(newSlice, 45)
// fmt.Println(slice)
// fmt.Println(newSlice)
// newSlice[0] = 15
// fmt.Println(slice)
// fmt.Println(newSlice)

// Example-3: When capcacity of new slice have no space due to defined capacity, underlying array gets
// changed due to append opera}tion
// slice := []int{10, 20, 30, 40, 50}
// newSlice := slice[1:3:3]
// fmt.Println(slice)
// fmt.Println(newSlice)
// newSlice = append(newSlice, 45)
// fmt.Println(slice)
// fmt.Println(newSlice)
// newSlice[0] = 15
// fmt.Println(slice)
// fmt.Println(newSlice)

// }
