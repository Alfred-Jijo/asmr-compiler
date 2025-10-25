package main

import (
	"fmt"
)

type __TokenType int

const (
	BYTE __TokenType = iota
	BYTED8
	LDV
	BEGIN
	END
	PLUS
	MINUS
	EQUAL_EQUAL
	GREATER_THAN
	LESS_THAN
	NOT
	READ
	PRINT
	DECL
	CALL
	BRANCH
	LIST
	END_LIST
)

var TokenID = map[__TokenType]string{
	//primitives
	BYTE:   "byte",
	BYTED8: "byted8",

	//variables
	LDV: "ldv",

	//operations
	BEGIN: "begin",
	END:   "end",

	PLUS:         "plus",
	MINUS:        "sulp",
	EQUAL_EQUAL:  "gcme",
	GREATER_THAN: "gcmg",
	LESS_THAN:    "gcml",
	NOT:          "gcmr",

	//io
	READ:  "udp",
	PRINT: "dmp",

	//functions
	DECL: "decl",
	CALL: "call",

	//selection
	BRANCH: "branc",

	//list
	LIST:     "list",
	END_LIST: "tsil",
}

type __Token struct {
	tokenType __TokenType
	lexeme    string
}

// --- NEW CODE STARTS HERE ---

// 1. Define the reverse map
var Keywords = make(map[string]__TokenType)

// 2. Populate the map automatically on program start
func init() {
	for tokenType, keyword := range TokenID {
		Keywords[keyword] = tokenType
	}
}

// 3. Example of how to use it
func token_main() {
	if DEBUG {
		// Simulate reading these words from a file
		words := []string{"byte", "ldv", "myVariable", "sulp"}

		for _, word := range words {
			// Look up the word in our new Keywords map
			tokenType, found := Keywords[word]

			if found {
				// It's a keyword!
				fmt.Printf("Read: '%s' -> Found keyword! Type: %d\n", word, tokenType)
			} else {
				// It's not a keyword
				fmt.Printf("Read: '%s' -> Not a keyword (maybe an identifier?)\n", word)
			}
		}
	}
}
