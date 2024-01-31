package main

// import (
// 	"fmt"
// 	"io"
// 	"strings"
// )

// func main() {
// 	myStringData := strings.NewReader("Akhand Agarwal")
// 	packet := make([]byte, 3)

// 	for {
// 		n, err := myStringData.Read(packet)
// 		fmt.Printf("%d bytes read , data : %s\n", n, packet[:n])

// 		// Handle error
// 		if err == io.EOF {
// 			fmt.Println("--EOF-REACHED--")
// 			break
// 		} else if err != nil {
// 			fmt.Println("Some Unexpected error occured")
// 			break
// 		}
// 	}
// }
