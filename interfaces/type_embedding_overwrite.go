package main

// type notifier interface {
// 	notify()
// }

// type user struct {
// 	name  string
// 	email string
// }

// func (u *user) notify() {
// 	fmt.Printf("sending user email to %s<%s>\n", u.name, u.email)
// }

// func (u *admin) notify() {
// 	fmt.Printf("sending admin email to %s<%s>\n", u.name, u.email)
// }

// type admin struct {
// 	user
// 	level string
// }

// func main() {
// 	ad := admin{
// 		user: user{
// 			name:  "Akhand Agarwal",
// 			email: "akhand.agarwal@gmail.com",
// 		},
// 		level: "super",
// 	}

// 	// ad.notify()
// 	// ad.user.notify()
// 	sendNotification(&ad)
// 	sendNotification(&ad.user)
// 	ad.notify()
// 	ad.user.notify()
// }

// func sendNotification(n notifier) {
// 	n.notify()
// }
