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

// func PrintDetails(i interface{}) {
// 	fmt.Printf("The type is %T\n", i)
// 	fmt.Printf("The value is %v\n", i)

// 	switch i.(type) {
// 	case Shape:
// 		forarea := i.(Shape)
// 		fmt.Printf("The type is Shape and area of the figure is %v\n", forarea.Area())
// 	case Object:
// 		forvolume := i.(Object)
// 		fmt.Printf("The type is Object and its volume is %v\n", forvolume.Volume())
// 	default:
// 		fmt.Printf("The type is Primitive type\n")
// 	}
// }

// func main() {
// 	c := Cube{8, 5, 3}
// 	var s Shape = c
// 	PrintDetails(s)
// 	var o Object = symmetric_cube{8}
// 	PrintDetails(o)
// 	PrintDetails("akhand")
// 	PrintDetails(56)
// }
