package main

import (
	"log"
	"os"
)

func main() {
	if err := os.Remove("testFile.txt"); err != nil {
		log.Fatal(err)
	}
}
