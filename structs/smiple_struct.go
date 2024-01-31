package main

import "fmt"

type employee struct {
	firstName, lastname string
	age                 int
	gender              bool
}

func main() {
	var akhand employee
	akhand.firstName = "akhand"
	akhand.lastname = "Agarwal"
	akhand.age = 21
	akhand.gender = false

	fmt.Printf("The details of the employee are %v", akhand)

	akhand2 := employee{
		firstName: "akhand",
		lastname:  "agarwal",
	}

	fmt.Println(akhand2)
}
