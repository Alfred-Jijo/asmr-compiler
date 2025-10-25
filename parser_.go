package main

import (
	"asmr-compiler/sound"
	"fmt"
	"log"
	"strconv"
	"os"
)

var mappings map[string]byte

func Parse(Tokens []__Token) {
	mappings = make(map[string]byte)
	var idx = 0
	var loop_body = 0
	var loop_end = 0
	for {
		if idx >= len(Tokens) {
			break
		}
		switch Tokens[idx].__tokenType {
		case LDV:
			sound.PlaySound("LDV")
			if DEBUG {
				fmt.Println("LDV")
			}
			var arg byte
			value, ok := mappings[Tokens[idx+3].lexeme]
			if ok {
				arg = value
			} else {
				val, _ := strconv.Atoi(Tokens[idx+3].lexeme)
				arg = byte(val)
			}
			mappings[Tokens[idx+2].lexeme] = arg
			idx += 4
			break
		case PLUS:
			sound.PlaySound("PLUS")
			if DEBUG {
				fmt.Println("PLUS")
			}
			var arg1 byte
			value, ok := mappings[Tokens[idx+2].lexeme]
			if ok {
				arg1 = value
			} else {
				val, _ := strconv.Atoi(Tokens[idx+2].lexeme)
				arg1 = byte(val)
			}

			var arg2 byte
			value, ok = mappings[Tokens[idx+3].lexeme]
			if ok {
				arg2 = value
			} else {
				val, _ := strconv.Atoi(Tokens[idx+3].lexeme)
				arg2 = byte(val)
			}
			mappings[Tokens[idx+1].lexeme] = arg1 + arg2
			idx += 4
			break
		case MINUS:
			sound.PlaySound("SULP")
			if DEBUG {
				fmt.Println("SULP or MINUS")
			}
			var arg1 byte
			value, ok := mappings[Tokens[idx+2].lexeme]
			if ok {
				arg1 = value
			} else {
				val, _ := strconv.Atoi(Tokens[idx+2].lexeme)
				arg1 = byte(val)
			}

			var arg2 byte
			value, ok = mappings[Tokens[idx+3].lexeme]
			if ok {
				arg2 = value
			} else {
				val, _ := strconv.Atoi(Tokens[idx+3].lexeme)
				arg2 = byte(val)
			}
			mappings[Tokens[idx+1].lexeme] = arg1 - arg2
			idx += 4
			break
		case READ:
			sound.PlaySound("UDP")
			fmt.Println("READ")

			if (Tokens[idx + 1].lexeme == "stdin") {
				var arg int
				fmt.Scan(&arg)
				barg := byte(arg)
				mappings[Tokens[idx+2].lexeme] = barg
			} else {
				file, _ := os.Open(Tokens[idx + 1].lexeme)
				buffer := make([]byte, 1)
				_, _ = file.Read(buffer)
				mappings[Tokens[idx+2].lexeme] = buffer[0]
			}

			idx += 3
			break
		case PRINT:
			sound.PlaySound("DMP")
			if DEBUG {
				fmt.Println("PRINT")
			}

			if (Tokens[idx + 1].lexeme == "stdout") {
				var arg byte
				value, ok := mappings[Tokens[idx+2].lexeme]
				if ok {
					arg = value
				} else {
					val, _ := strconv.Atoi(Tokens[idx+2].lexeme)
					arg = byte(val)
				}
				fmt.Printf("%d\n", arg)
			} else {
				var out byte;
				value, ok := mappings[Tokens[idx+2].lexeme]
				if ok {
					out = value
				} else {
					val, _ := strconv.Atoi(Tokens[idx+2].lexeme)
					out = byte(val)
				}

				name := Tokens[idx + 1].lexeme
				file, _ := os.OpenFile(name, os.O_RDWR | os.O_CREATE, 0644)
				_, _ = file.Write([]byte{out})
			}
			idx += 3
			break
		case EQUAL_EQUAL:
			sound.PlaySound("GCM")
			if DEBUG {
				fmt.Println("GCME")
			}
			var arg1 byte
			value, ok := mappings[Tokens[idx+1].lexeme]
			if ok {
				arg1 = value
			} else {
				val, _ := strconv.Atoi(Tokens[idx+1].lexeme)
				arg1 = byte(val)
			}

			var arg2 byte
			value, ok = mappings[Tokens[idx+2].lexeme]
			if ok {
				arg2 = value
			} else {
				val, _ := strconv.Atoi(Tokens[idx+2].lexeme)
				arg2 = byte(val)
			}

			if arg1 == arg2 {
				idx += 3
			} else {
				var offset = 3
				for Tokens[idx+offset].__tokenType != ELSE {
					offset += 1
				}
				idx += offset + 1
			}
		case LESS_THAN:
			sound.PlaySound("GCM")
			if DEBUG {
				fmt.Println("GCML")
			}
			var arg1 byte
			value, ok := mappings[Tokens[idx+1].lexeme]
			if ok {
				arg1 = value
			} else {
				val, _ := strconv.Atoi(Tokens[idx+1].lexeme)
				arg1 = byte(val)
			}

			var arg2 byte
			value, ok = mappings[Tokens[idx+2].lexeme]
			if ok {
				arg2 = value
			} else {
				val, _ := strconv.Atoi(Tokens[idx+2].lexeme)
				arg2 = byte(val)
			}

			if arg1 < arg2 {
				idx += 3
			} else {
				var offset = 3
				for Tokens[idx+offset].__tokenType != ELSE {
					offset += 1
				}
				idx += offset + 1
			}
		case GREATER_THAN:
			sound.PlaySound("GCM")
			if DEBUG {
				fmt.Println("GCMG")
			}
			var arg1 byte
			value, ok := mappings[Tokens[idx+1].lexeme]
			if ok {
				arg1 = value
			} else {
				val, _ := strconv.Atoi(Tokens[idx+1].lexeme)
				arg1 = byte(val)
			}

			var arg2 byte
			value, ok = mappings[Tokens[idx+2].lexeme]
			if ok {
				arg2 = value
			} else {
				val, _ := strconv.Atoi(Tokens[idx+2].lexeme)
				arg2 = byte(val)
			}

			if arg1 > arg2 {
				idx += 3
			} else {
				var offset = 3
				for Tokens[idx+offset].__tokenType != ELSE {
					offset += 1
				}
				idx += offset + 1
			}
		case ELSE:
			sound.PlaySound("ALT")
			var offset = 1
			for Tokens[idx+offset].__tokenType != END {
				offset += 1
			}
			idx += offset + 1

		case END:
			sound.PlaySound("END")
			idx += 1
			break

		case LOOP:
			loop_body = idx + 1
			loop_end = loop_body
			for Tokens[loop_end].__tokenType != LOOP_CLOSE {
				loop_end += 1
			}
			loop_end += 1
			idx += 1

		case BREAK:
			idx = loop_end

		case LOOP_CLOSE:
			idx = loop_body

		default:
			sound.PlaySound("ERROR")
			log.Fatal("Unknown TokenType")
		}
	}
}
