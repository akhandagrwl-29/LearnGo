package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// // Sometimes, we do not want to encode a field value as it is but to provide a custom value for marshaling. This can be achieved by implementing json.Marshaler or encoding.TextMarshaler interface.

// type Profile struct {
// 	Username  string
// 	Followers int
// }

// type Age int

// func (p Profile) MarshalJSON() ([]byte, error) {
// 	return []byte(fmt.Sprintf(`{"f_count" : "%d"}`, p.Followers)), nil
// }

// func (a Age) MarshalText() ([]byte, error) {
// 	return []byte(fmt.Sprintf(`{"age": %d}`, int(a))), nil
// }

// type Student struct {
// 	FirstName, lastName string
// 	Age                 Age
// 	Profile             Profile
// }

// func main() {
// 	john := &Student{
// 		FirstName: "Akhand",
// 		lastName:  "Agarwal",
// 		Age:       22,
// 		Profile: Profile{
// 			Username:  "akhand_agarwal",
// 			Followers: 1200,
// 		},
// 	}

// 	johnJSON, _ := json.MarshalIndent(john, "", " ")

// 	fmt.Println(string(johnJSON))
// }
