package main

import (
	"fmt"
	"strings"
	"unicode"
)

var Keywords = make(map[string]__TokenType)

func Lex(data []byte) []__Token {
	line := string(data)
	lines := splitMDelim(line, " \n")

	/*if DEBUG {
		for _, line := range lines {
			fmt.Println(line)
		}
	}*/

	populateRmap()

	var tokens []__Token
	for _, line := range lines {
		tokenType, found := Keywords[line]
		if found {
			tokens = append(tokens, __Token{__tokenType: tokenType, lexeme: line})
		} else {
			tokens = append(tokens, __Token{__tokenType: NONE, lexeme: line})
		}
		if DEBUG {
			for _, token := range tokens {
				fmt.Println(token.lexeme)
				printEnumVal(token.__tokenType)
			}
		}
	}
	fmt.Println(tokens)
	return tokens
}

func populateRmap() {
	for tokenType, keyword := range TokenID {
		Keywords[keyword] = tokenType
	}
}

func splitMDelim(s string, delimiters string) []string {
	delims := make(map[rune]bool)
	for _, d := range delimiters {
		delims[d] = true
	}

	// FieldsFunc splits when the function returns true
	return strings.FieldsFunc(s, func(r rune) bool {
		return delims[r] || unicode.IsSpace(r)
	})
}
