package pointersVsValues

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	Length float64
	Breath float64
}

//func (s Rect) Area() float64 {
//	return s.length * s.breath
//}
//
//func (s *Rect) Perimeter() float64 {
//	return 2 * (s.length + s.breath)
//}
//
//func main() {
//
//	var s Shape = &Rect{8, 9}
//	fmt.Printf("The area of the figure is %v\n", s.Area())
//	fmt.Printf("The perimter of the figure is %v\n", s.Perimeter())
//}
