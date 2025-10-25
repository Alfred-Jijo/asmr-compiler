package main

import (
	"fmt"
	"strings"
)

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
