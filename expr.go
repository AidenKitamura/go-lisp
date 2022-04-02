package main

import "strconv"

const (
	ADD = iota
	MINUS
	MULTIPLY
	DIVIDE
	INT
	UNKNOWN
)

type EXPR struct {
	expr     int
	item     interface{}
	operands []*EXPR
}

func (e *EXPR) EVAL() interface{} {
	if e.expr == INT {
		return e.item
	}
	res := 0
	switch e.expr {
	case ADD:
		for _, op := range e.operands {
			res += op.EVAL().(int)
		}
	case MINUS:
		for _, op := range e.operands {
			res -= op.EVAL().(int)
		}
	case MULTIPLY:
		res = 1
		for _, op := range e.operands {
			res *= op.EVAL().(int)
		}
	case DIVIDE:
		res = e.operands[0].EVAL().(int) * e.operands[0].EVAL().(int)
		for _, op := range e.operands {
			res -= op.EVAL().(int)
		}
	default:
		panic("error, non evalable code")
	}
	return res

}

func GEN_TOKENS(s string) []string {
	tokens := make([]string, 0)
	ptr := 0
	state := 0
	for i := 0; i < len(s); i++ {
		switch state {
		case 0:
			// Starting state
			// Ignore until non space
			// If find symbols
			// Add to tokens
			// If find numericals
			// Start counting
			switch s[i] {
			case ' ', '\n':
				// Skip space
				continue
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				// Find numbers
				// Start counting
				ptr = i
				state = 1
			default:
				// Symbol
				// Simply append
				tokens = append(tokens, string(s[i]))
			}

		case 1:
			// Numerical Values
			// Count until end
			switch s[i] {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				// Numericals Continue
				continue
			case ' ', '\n':
				// End of Numericals
				// Append numerical as token
				// Go back to state 0
				tokens = append(tokens, s[ptr:i])
				state = 0
			default:
				// Possible symbol
				// Append both numemricals and symbols
				tokens = append(tokens, s[ptr:i], string(s[i]))
				state = 0
			}
		}
	}

	return tokens
}

func S_GEN_EXPRS(s string) *EXPR {
	tokens := GEN_TOKENS(s)
	exp := GEN_EXPRS(tokens)
	return exp
}

func GEN_EXPRS(tokens []string) *EXPR {
	ex := EXPR{}
	state := 0
	for _, token := range tokens {
		switch state {
		case 0:
			// waiting for start symbol
			// Initial state
			// only appear once
			switch token {
			case "(":
				state = 1
			default:
				// Error detected
				// Invalid start symbol
				panic("invalid starting symbol detected, abort")
			}

		case 1:
			// Looking for expression type
			// Create expressions
			// and operands list
			// Then enter state 2 to deal
			// with possible operands
			TYPE := DEC_TYPE(token)
			ex.expr = TYPE
			if TYPE == INT {
				ex.item, _ = strconv.Atoi(token)
			} else {
				ex.item = token
			}
			ex.operands = make([]*EXPR, 0)
			state = 2

		case 2:
			switch token {
			case ")":
				// Already parsed, no more operands
				break
			case "(":
				// Need to recursively create operand
				// Need extra things to do
			default:
				// Add operands
				p_result, _ := strconv.Atoi(token)
				ex.operands = append(ex.operands, &EXPR{DEC_TYPE(token), p_result, nil})
			}
		}
	}
	return &ex
}

func DEC_TYPE(s string) int {
	switch s {
	case "+":
		return ADD
	case "-":
		return MINUS
	case "*":
		return MULTIPLY
	case "/":
		return DIVIDE
	default:
		return INT
	}

}
