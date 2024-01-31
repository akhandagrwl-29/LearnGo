package main

// import (
// 	"fmt"
// 	"io"
// )

// type MyStringData struct {
// 	str     string
// 	readIdx int
// }

// func (myStringData *MyStringData) Read(p []byte) (n int, err error) {
// 	strBytes := []byte(myStringData.str)

// 	if myStringData.readIdx >= len(strBytes) {
// 		return 0, io.EOF
// 	}

// 	nextReadIdx := myStringData.readIdx + len(p)

// 	if nextReadIdx >= len(strBytes) {
// 		nextReadIdx = len(strBytes)
// 		err = io.EOF
// 	}

// 	nextBytes := strBytes[myStringData.readIdx:nextReadIdx]
// 	n = len(nextBytes)

// 	copy(p, nextBytes)
// 	myStringData.readIdx = nextReadIdx

// 	return
// }

// func main() {
// 	myStringData := &MyStringData{
// 		str: "Akhand Agarwal",
// 	}

// 	packet := make([]byte, 3)

// 	for {
// 		n, err := myStringData.Read(packet)
// 		fmt.Printf("%d bytes were read, data : %s\n", n, packet[:n])

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
