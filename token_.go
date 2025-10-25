package main

type __TokenType int

const (
	BYTE __TokenType = iota
	BYTED8
	LDV
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
	ELSE
	END
	NONE
	LOOP
	BREAK
	LOOP_CLOSE
)

var TokenID = map[__TokenType]string{
	//primitives
	BYTE:   "byte",
	BYTED8: "byted8",

	//variables
	LDV: "ldv",


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
	ELSE: "alt",
	END:   "end",

	//loop
	LOOP: "spin",
	BREAK: "crash",
	LOOP_CLOSE: "flight",
}

type __Token struct {
	__tokenType __TokenType
	lexeme      string
}

func printEnumVal(toktype __TokenType) {
}
