package main

import (
	"fmt"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func ReadFile(name string) []byte {
	data, err := os.ReadFile(name)
	check(err)
	return data
}

func main() {
	name := os.args[1]
}
