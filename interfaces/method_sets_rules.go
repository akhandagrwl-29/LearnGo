package main

// type notifier interface {
// 	notify()
// }

// type user struct {
// 	name  string
// 	email string
// }

// // notify implements a method with pointer receiver
// func (u *user) notify() {
// 	fmt.Printf("Sending user email to %s and %s", u.name, u.email)
// }

// func main() {
// 	u := user{"akhand", "akhand@gmail.com"}
// 	sendNotification(u)
// }

// func sendNotification(n notifier) {
// 	n.notify()
// }

// This program won't compile because of violating method set rules.
// The type of receiever is used to determine whether a method is associated with value , pointer or both

// Rules
// T         (t T)
// *T        (t T) and (t *T)

// Value of type T only has methods declared that have value receiver as part of it's method set
// Pointer of type T has methods declared that have value receiver as well as pointer receiever
//  as part of it's method set

// From the perspective of reciever - Easy to remember

// (t T)    T and *T
// (t *T)   *T
