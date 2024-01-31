package main

// type Object interface {
// 	Volume() float64
// }
// type Shape interface {
// 	Area() float64
// }

// type Temp interface {
// 	Height() float64
// }

// type Cube struct {
// 	length float64
// 	breath float64
// 	height float64
// }

// type symmetric_cube struct {
// 	length float64
// }

// func (s symmetric_cube) Volume() float64 {
// 	return s.length * s.length * s.length
// }

// func (s Cube) Area() float64 {
// 	return s.length * s.breath
// }

// func main() {
// 	var s Shape
// 	s = Cube{8, 5, 3}
// 	fmt.Printf("The area of the figure is %v\n", s.Area())
// 	_, ok := s.(Temp)
// 	if ok == true {
// 		fmt.Printf("Different interface can also be type asserted")
// 	} else {
// 		fmt.Printf("Different interface can not be type asserted\n")
// 	}
// }
