package main

// import (
// 	"fmt"
// 	"io"
// 	"strings"
// )

// func main() {
// 	mainSrc := strings.NewReader("Akhand Agarwal")
// 	src := io.LimitReader(mainSrc, 6)

// 	buf := make([]byte, 3)
// 	for {
// 		n, err := src.Read(buf)
// 		fmt.Printf("%d bytes read , data: %s\n", n, buf[:n])

// 		if err == io.EOF {
// 			fmt.Println("End of file reached")
// 			break
// 		} else if err != nil {
// 			fmt.Printf("%v error occured\n", err)
// 			break
// 		}
// 	}
// }
