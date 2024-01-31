package main

// import "fmt"

// type Product interface {
// 	GetName() string
// }

// type Factory interface {
// 	CreateProduct() Product
// }

// type ConcreteMacProduct struct {
// 	name string
// }

// type ConcreteWinProduct struct {
// 	name string
// }

// type ConcreteMacFactory struct{}

// func (p *ConcreteMacProduct) GetName() string {
// 	return p.name
// }

// func (p *ConcreteMacFactory) CreateProduct() Product {
// 	return &ConcreteMacProduct{
// 		name: "Macbook Pro",
// 	}
// }

// type ConcreteWinFactory struct{}

// func (p *ConcreteWinProduct) GetName() string {
// 	return p.name
// }

// func (p *ConcreteWinFactory) CreateProduct() Product {
// 	return &ConcreteWinProduct{
// 		name: "Windows 11",
// 	}
// }

// func main() {
// 	macFactory := &ConcreteMacFactory{}
// 	product := macFactory.CreateProduct()
// 	fmt.Println(product.GetName())

// 	winFactory := &ConcreteWinFactory{}
// 	product = winFactory.CreateProduct()
// 	fmt.Println(product.GetName())
// }
