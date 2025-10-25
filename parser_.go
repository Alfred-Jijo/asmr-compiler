package main

var mappings map[string]byte

func ParseData(Tokens []__TokenType) {

	var idx = 0;
	loop {
		if idx >= Tokens.len {
			break
		}
		switch(Tokens[idx]) 
		case LDV:
			fmt.Println("LDV")
			mappings[Tokens[idx + 2]], err := byte(strconv.ParseUint(Tokens[idx + 3]))
			idx += 4	
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
		case READ:
			fmt.Println("READ")
			break
		case PRINT:
			fmt.Println("PRINT")
			break
		default:
			fmt.Println("Unknown TokenType")
		}
	}
	case LDV:
		fmt.Println("LDV")
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
	case READ:
		fmt.Println("READ")
		break
	case PRINT:
		fmt.Println("PRINT")
		break
	default:
		fmt.Println("Unknown TokenType")
	}
}
