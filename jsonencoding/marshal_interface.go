package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Iprofile interface {
// 	Follow()
// }

// type Profile struct {
// 	Username  string
// 	Followers int
// }

// func (p *Profile) Follow() {
// 	p.Followers++
// }

// type Student struct {
// 	FirstName, lastName string
// 	Age                 int
// 	Primary             Iprofile
// 	Secondary           Iprofile
// }

// func main() {
// 	john := &Student{
// 		FirstName: "John",
// 		lastName:  "Doe",
// 		Age:       22,
// 		Primary: &Profile{
// 			Username:  "Jondow32",
// 			Followers: 1975,
// 		},
// 	}
// 	john.Primary.Follow()
// 	johnJSON, _ := json.MarshalIndent(john, "", " ")

// 	fmt.Println(string(johnJSON))
// }
