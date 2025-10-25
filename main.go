package main

import (
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

func Lex(data []byte){

}

func ParseData(data []byte) {

}

func main() {
	name := os.Args[1]
	data := ReadFile(name)
	os.Stdout.Write(data)
	//ParseData(data)
}
