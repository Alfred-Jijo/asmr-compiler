package main

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
