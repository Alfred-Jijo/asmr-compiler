package main

import (
	"fmt"
	"strings"
)

func Lex(data []byte) {
	line := string(data)
	lines := strings.Split(line, " ")

	tokens []_token
	for _, line := range lines{
		tokens.append(_token{type:  ,lexeme: _})
	}

	if DEBUG {
		for _, line := range lines {
			fmt.Println(line)
		}
	}
}
