package main

import "fmt"

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
	NONE
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
	__tokenType __TokenType
	lexeme      string
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
