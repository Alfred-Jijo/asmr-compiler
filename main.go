package main

import (
	"fmt"
	"log"
	"os"
)

var (
	DEBUG = false
)

func check(e error, printFlag string) {
	if e != nil {
		if printFlag == "log" {
			log.Fatal(e)
		} else {
			panic(e)
		}
	}
}

func ReadFile(name string) []byte {
	data, err := os.ReadFile(name)
	check(err, "log")
	return data
}

func main() {
	name := os.Args[1]
	if len(os.Args) > 2 {
		DEBUG = true
	}
	if DEBUG {
		fmt.Printf("Reading file %s\n", name)
	}
	token_main()
	// data := ReadFile(name)

	// _, err := os.Stdout.Write(data)
	// check(err, "log")

	//Lex(data)
}
