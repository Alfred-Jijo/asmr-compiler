package main

import (
	"fmt"
	"log"
	"os"
	"strings"
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

func Lex(data []byte) {
	line := string(data)
	lines := strings.Split(line, " ")
	
	tokens []__Token;
	for _, line := range lines{
		tokens.append(__Token{tokenType:  ,lexeme: _})
	}

	if DEBUG {
		for _, line := range lines {
			fmt.Println(line)
		}
	}
}

func ParseData(data []byte) {
}

func main() {
	name := os.Args[1]
	if len(os.Args) > 2 {
		DEBUG = true
	}
	data := ReadFile(name)

	// _, err := os.Stdout.Write(data)
	// check(err, "log")

	Lex(data)
}
