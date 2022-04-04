// expr.go includes definition for EXPR, which
// is the minimal unit in LISP. In this imple-
// mentation, every single unit except for '('
// and ')' are considered an expression.

package parser

import (
	"fmt"
	"strconv"
)

// The expression might be several types, which
// are numericals, lists or operations.
const (
	ADD = iota
	MINUS
	MULTIPLY
	DIVIDE
	INT
	FLOAT
	LIST
	UNKNOWN
	ERROR
)

type EXPR struct {
	op       int     // Type
	item     string  // Value
	operands []*EXPR // Possible Operands
}

// Eval function provides evaluation function
// for lisp expressions, currently the defini-
// tion is like this:
//
// If the EXPR is numerical operation, the num-
// ber of operands must be >= 2, otherwise the
// interpreter stop evaluating and throw an er-
// ror
//
// If the EXPR is numerical, the value and its
// type is returned.
//
// If the EXPR is list, the string of the oper-
// ands are returned
func EVAL(e *EXPR) (TYPE int, VALUE interface{}) {
	switch e.op {
	case ADD:
		return EVAL_ADD(e)
	case MINUS:
		return EVAL_MINUS(e)
	case MULTIPLY:
		return EVAL_MULTIPLY(e)
	case DIVIDE:
		return EVAL_DIVIDE(e)
	case INT:
		operand, err := strconv.Atoi(e.item)
		if err == nil {
			return INT, operand
		} else {
			return ERROR, err
		}
	case FLOAT:
		operand, err := strconv.ParseFloat(e.item, 64)
		if err == nil {
			return FLOAT, operand
		} else {
			return ERROR, err
		}
	case LIST:
		// Need to change
		return LIST, "list"
	default:
		return ERROR, fmt.Errorf("unknown expression")
	}
}

// ADD function deals with add expressions. We
// first check if there are any floats in the
// operands, and if there is, we convert every-
// thing to float and return float result
// ***Allows single operands***
func EVAL_ADD(e *EXPR) (TYPE int, VALUE interface{}) {
	// Check if any of the operands are float type
	// If yes, then type cast everything to float
	// In order to boost performance, use a slice
	// to store pre-computed results, if any
	NEED_FLOAT := false
	PRE_RESULT := make([]interface{}, 0)
	for _, ex := range e.operands {
		switch ex.op {
		case INT:
			// Add integer to precomputed list
			// If error, return error
			T, V := EVAL(ex)
			if T == ERROR {
				return ERROR, V
			} else {
				PRE_RESULT = append(PRE_RESULT, V)
			}

		case FLOAT:
			// Add float to precomputed list
			// If error, return error
			NEED_FLOAT = true
			T, V := EVAL(ex)
			if T == ERROR {
				return ERROR, V
			} else {
				PRE_RESULT = append(PRE_RESULT, V)
			}

		case ADD:
			// If find a sub expression, evaluate
			// it and see if it is float
			T, V := EVAL_ADD(ex)
			switch T {
			case INT:
				PRE_RESULT = append(PRE_RESULT, V)
			case FLOAT:
				PRE_RESULT = append(PRE_RESULT, V)
				NEED_FLOAT = true
			default:
				return ERROR, fmt.Errorf("unknown value type: %d", T)
			}

		case MINUS:
			// If find a sub expression, evaluate
			// it and see if it is float
			T, V := EVAL_MINUS(ex)
			switch T {
			case INT:
				PRE_RESULT = append(PRE_RESULT, V)
			case FLOAT:
				PRE_RESULT = append(PRE_RESULT, V)
				NEED_FLOAT = true
			default:
				return ERROR, fmt.Errorf("unknown value type: %d", T)
			}

		case MULTIPLY:
			// If find a sub expression, evaluate
			// it and see if it is float
			T, V := EVAL_MULTIPLY(ex)
			switch T {
			case INT:
				PRE_RESULT = append(PRE_RESULT, V)
			case FLOAT:
				PRE_RESULT = append(PRE_RESULT, V)
				NEED_FLOAT = true
			default:
				return ERROR, fmt.Errorf("unknown value type: %d", T)
			}

		case DIVIDE:
			// If find a sub expression, evaluate
			// it and see if it is float
			T, V := EVAL_DIVIDE(ex)
			switch T {
			case INT:
				PRE_RESULT = append(PRE_RESULT, V)
			case FLOAT:
				PRE_RESULT = append(PRE_RESULT, V)
				NEED_FLOAT = true
			default:
				return ERROR, fmt.Errorf("unknown value type: %d", T)
			}

		default:
			return ERROR, fmt.Errorf("unknown non-numerical operand type: %d", ex.op)
		}
	}

	if NEED_FLOAT {
		// type case to float
		// then return value
		res := 0.0
		for _, num := range PRE_RESULT {
			// Check if is float64 type
			operand, ok := num.(float64)
			if ok {
				res += operand
			} else {
				res += float64(num.(int))
			}
		}
		return FLOAT, res
	} else {
		// All ints
		res := 0
		for _, num := range PRE_RESULT {
			res += num.(int)
		}
		return INT, res
	}
}

// MINUS function deals with minus expressions.
// We first check if there are any floats in the
// operands, and if there is, we convert every-
// thing to float and return float result
// ***Allows single operands***
func EVAL_MINUS(e *EXPR) (TYPE int, VALUE interface{}) {
	// Check if any of the operands are float type
	// If yes, then type cast everything to float
	// In order to boost performance, use a slice
	// to store pre-computed results, if any
	NEED_FLOAT := false
	PRE_RESULT := make([]interface{}, 0)
	for _, ex := range e.operands {
		switch ex.op {
		case INT:
			// Add integer to precomputed list
			// If error, return error
			T, V := EVAL(ex)
			if T == ERROR {
				return ERROR, V
			} else {
				PRE_RESULT = append(PRE_RESULT, V)
			}

		case FLOAT:
			// Add float to precomputed list
			// If error, return error
			NEED_FLOAT = true
			T, V := EVAL(ex)
			if T == ERROR {
				return ERROR, V
			} else {
				PRE_RESULT = append(PRE_RESULT, V)
			}

		case ADD:
			// If find a sub expression, evaluate
			// it and see if it is float
			T, V := EVAL_ADD(ex)
			switch T {
			case INT:
				PRE_RESULT = append(PRE_RESULT, V)
			case FLOAT:
				PRE_RESULT = append(PRE_RESULT, V)
				NEED_FLOAT = true
			default:
				return ERROR, fmt.Errorf("unknown value type: %d", T)
			}

		case MINUS:
			// If find a sub expression, evaluate
			// it and see if it is float
			T, V := EVAL_MINUS(ex)
			switch T {
			case INT:
				PRE_RESULT = append(PRE_RESULT, V)
			case FLOAT:
				PRE_RESULT = append(PRE_RESULT, V)
				NEED_FLOAT = true
			default:
				return ERROR, fmt.Errorf("unknown value type: %d", T)
			}

		case MULTIPLY:
			// If find a sub expression, evaluate
			// it and see if it is float
			T, V := EVAL_MULTIPLY(ex)
			switch T {
			case INT:
				PRE_RESULT = append(PRE_RESULT, V)
			case FLOAT:
				PRE_RESULT = append(PRE_RESULT, V)
				NEED_FLOAT = true
			default:
				return ERROR, fmt.Errorf("unknown value type: %d", T)
			}

		case DIVIDE:
			// If find a sub expression, evaluate
			// it and see if it is float
			T, V := EVAL_DIVIDE(ex)
			switch T {
			case INT:
				PRE_RESULT = append(PRE_RESULT, V)
			case FLOAT:
				PRE_RESULT = append(PRE_RESULT, V)
				NEED_FLOAT = true
			default:
				return ERROR, fmt.Errorf("unknown value type: %d", T)
			}

		default:
			return ERROR, fmt.Errorf("unknown non-numerical operand type: %d", ex.op)
		}
	}

	if NEED_FLOAT {
		// type case to float
		// then return value
		if len(PRE_RESULT) == 1 {
			// Must be float, need not check
			return FLOAT, -PRE_RESULT[0].(float64)
		}

		// If more than one operands
		// Add first operand
		var res float64
		operand, ok := PRE_RESULT[0].(float64)
		if ok {
			res += operand
		} else {
			res += float64(PRE_RESULT[0].(int))
		}
		for index := 1; index < len(PRE_RESULT); index++ {
			operand, ok = PRE_RESULT[index].(float64)
			if ok {
				res -= operand
			} else {
				res -= float64(PRE_RESULT[index].(int))
			}
		}
		return FLOAT, res
	} else {
		// All ints, need not check
		res := 0
		if len(PRE_RESULT) == 1 {
			return INT, -PRE_RESULT[0].(int)
		} else {
			res += PRE_RESULT[0].(int)
			for index := 1; index < len(PRE_RESULT); index++ {
				res -= PRE_RESULT[index].(int)
			}
		}
		return INT, res
	}
}

// ADD function deals with add expressions. We
// first check if there are any floats in the
// operands, and if there is, we convert every-
// thing to float and return float result
// ***Allow Single Operands***
func EVAL_MULTIPLY(e *EXPR) (TYPE int, VALUE interface{}) {
	return ERROR, "multi error"
}

// ADD function deals with add expressions. We
// first check if there are any floats in the
// operands, and if there is, we convert every-
// thing to float and return float result
// ***Allow Single Operands***
func EVAL_DIVIDE(e *EXPR) (TYPE int, VALUE interface{}) {
	return ERROR, "divide error"
}
