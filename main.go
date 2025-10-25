package main

import (
	"fmt"
	"os"
)

func CreateFile() {
	fmt.Println("Writing file")
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	length, err := file.WriteString("Welcome to Golang. This demonstrates file operations.")
	if err != nil {
		panic(err)
	}

	fmt.Printf("File name: %s\n", file.Name())
	fmt.Printf("File length: %d bytes\n", length)
}
func main() {

}
