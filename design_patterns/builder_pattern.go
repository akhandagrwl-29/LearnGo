package main

// import "fmt"

// type Pizza struct {
// 	dough    string
// 	sauce    string
// 	cheese   string
// 	toppings []string
// }

// type IpizzaBuilder interface {
// 	SetDough(string) IpizzaBuilder
// 	SetSauce(string) IpizzaBuilder
// 	SetCheese(string) IpizzaBuilder
// 	SetToppings([]string) IpizzaBuilder
// 	Build() *Pizza
// }

// type ConcretePizzaBuilder struct {
// 	pizza *Pizza
// }

// func NewConcretePizzaBuilder() *ConcretePizzaBuilder {
// 	return &ConcretePizzaBuilder{pizza: &Pizza{}}
// }

// func (cpb *ConcretePizzaBuilder) SetDough(dough string) IpizzaBuilder {
// 	cpb.pizza.dough = dough
// 	return cpb
// }

// func (cpb *ConcretePizzaBuilder) SetSauce(sauce string) IpizzaBuilder {
// 	cpb.pizza.sauce = sauce
// 	return cpb
// }

// func (cpb *ConcretePizzaBuilder) SetCheese(cheese string) IpizzaBuilder {
// 	cpb.pizza.cheese = cheese
// 	return cpb
// }

// func (cpb *ConcretePizzaBuilder) SetToppings(toppings []string) IpizzaBuilder {
// 	cpb.pizza.toppings = toppings
// 	return cpb
// }

// func (cpb *ConcretePizzaBuilder) Build() *Pizza {
// 	return cpb.pizza
// }

// type Director struct {
// 	builder IpizzaBuilder
// }

// func NewDirector(builder IpizzaBuilder) *Director {
// 	return &Director{builder: builder}
// }

// func (d *Director) Construct() *Pizza {
// 	return d.builder.SetDough("Thin Crust").SetSauce("Tomato").SetCheese("Mozzarella").SetToppings([]string{"Mushrooms", "Olives", "Onions"}).Build()
// }

// func main() {
// 	builder := NewConcretePizzaBuilder()
// 	director := NewDirector(builder)
// 	pizza := director.Construct()

// 	fmt.Println(pizza)
// }
