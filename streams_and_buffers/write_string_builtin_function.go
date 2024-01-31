package main

// import (
// 	"fmt"
// 	"io"
// )

// type SampleStore struct {
// 	data []byte
// }

// func (ss *SampleStore) Write(p []byte) (n int, err error) {
// 	// check if 10 bytes have already been written
// 	if len(ss.data) == 10 {
// 		return 0, io.EOF
// 	}

// 	remainCap := 10 - len(ss.data)
// 	writeLength := len(p)

// 	if remainCap <= writeLength {
// 		writeLength = remainCap
// 		err = io.EOF
// 	}

// 	ss.data = append(ss.data, p[:writeLength]...)
// 	n = writeLength
// 	return
// }

// func main() {
// 	ss := &SampleStore{}

// 	// Write 1 : "Hello"
// 	bytesWritten1, err1 := io.WriteString(ss, "Hello")
// 	fmt.Printf("Bytes Written: %d, error: %v\n", bytesWritten1, err1)

// 	// Write 1 : "Akhand"
// 	bytesWritten2, err2 := io.WriteString(ss, "Akhand")
// 	fmt.Printf("Bytes Written: %d, error: %v\n", bytesWritten2, err2)

// 	// Write 1 : "Agarwal"
// 	bytesWritten3, err3 := io.WriteString(ss, "Agarwal")
// 	fmt.Printf("Bytes Written: %d, error: %v\n", bytesWritten3, err3)

// }
