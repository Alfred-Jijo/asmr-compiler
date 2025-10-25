package main

import (
	"fmt"
	"strings"
	"unicode"
)

var Keywords = make(map[string]__TokenType)

func Lex(data []byte) {
	line := string(data)
	// lines := strings.Split(line, " ")
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
}

func printEnumVal(toktype __TokenType) {
	switch toktype {
	case BYTE:
		fmt.Println("BYTE")
		break
	case BYTED8:
		fmt.Println("BYTED8")
		break
	case LDV:
		fmt.Println("LDV")
		break
	case BEGIN:
		fmt.Println("BEGIN")
		break
	case END:
		fmt.Println("END")
		break
	case PLUS:
		fmt.Println("PLUS")
		break
	case MINUS:
		fmt.Println("SULP or MINUS")
		break
	case EQUAL_EQUAL:
		fmt.Println("EQUAL_EQUAL")
		break
	case GREATER_THAN:
		fmt.Println("GREATER_THAN")
		break
	case LESS_THAN:
		fmt.Println("LESS_THAN")
		break
	case NOT:
		fmt.Println("NOT")
		break
	case READ:
		fmt.Println("READ")
		break
	case PRINT:
		fmt.Println("PRINT")
		break
	case DECL:
		fmt.Println("DECL")
		break
	case CALL:
		fmt.Println("CALL")
		break
	case BRANCH:
		fmt.Println("BRANCH")
		break
	case LIST:
		fmt.Println("LIST")
		break
	case END_LIST:
		fmt.Println("END_LIST")
		break
	case NONE:
		fmt.Println("NONE")
		break
	default:
		fmt.Println("Unknown TokenType")
	}
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
