package main

import (
	"fmt"
	"os"
)

func ReadFile(name string) []bytes {
	data, err != os.ReadFile(name) 
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func main() {
	name = os.args[1]
}
