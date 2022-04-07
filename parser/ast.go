// ast includes functions to help build an
// abstract syntax tree. Errors are checked
// here and if parsed successfully, the pa-
// rsed result is returned to the user

package parser

import (
	"fmt"
)

// GEN_AST parses the tokens into an abstr-
// act syntax tree and if parsed correctly,
// the root expression is returned. If not,
// nil is returned with an error.
func GEN_AST(tokens []string) (*EXPR, error) {
	// edge case when no tokens are passed in
	if len(tokens) == 0 {
		return nil, fmt.Errorf("creating ast from nil tokens")
	}

	// we check if the parentheses are met during
	// parsing. If not, nil and error are returned
	Q := ASTQUEUEConstructer(tokens)
	root, err := PAR_AST(Q)
	if err != nil {
		return nil, err
	}
	return root, nil
}

// PAR_AST receives an ASTQUEUE and deals
// with the actual parsing job. All exprs
// and the operands are checked validity
// and if there are semantic errors, the
// parsing job stops and return an error.
// The errors are returned by concatnation
func PAR_AST(q *ASTQUEUE) (*EXPR, error) {
	// Get next token
	sign, err := q.Pop()

	// Check valid begin token
	if sign != BEGIN_EXPR || err != nil {
		return nil, fmt.Errorf("expecting begin token (, got %s, followed by: %s", sign, err)
	}

	// We parse until we find a new begin
	// token or we reached an end and by
	// default we set the expr to be bool
	// with value false, in case of direct
	// ends.
	ex := &EXPR{op: BOOLEAN, item: "0", operands: make([]*EXPR, 0)}
	for err == nil {
		// Check next element first. If we have
		// a begin token, start a new parsing f
		// unction, otherwise parse accordingly
		sign, err = q.Top()
		if err != nil {
			break
		}

		switch sign {
		case BEGIN_EXPR:
			// parse nested expressions
			nestedEx, err := PAR_AST(q)
			if err != nil {
				break
			}
			ex.operands = append(ex.operands, nestedEx)

		case END_EXPR:
			// Send the expression back
			switch ex.op {
			case ADD, MINUS, MULTIPLY, DIVIDE:
				// Check if got no operands. These shold
				// have operands.
				if len(ex.operands) == 0 {
					return nil, fmt.Errorf("non-nil operation %s with nil operands", sign)
				}
				fallthrough

			default:
				q.Pop()
				return ex, nil
			}

		case ADD_EXPR:
			// can only have one operation expression
			// that is, at the beginning. If the first
			// element encountered is not an operation
			// then the rest must not be operations
			if ex.op != BOOLEAN {
				return nil, fmt.Errorf("duplicate operation sign: %s and %s", ADD_EXPR, ex.item)
			}

			// if no error, set the expression to be
			// add and the item to be "+". After that
			// we parse the rest items. However, we
			// should check it there are no operands
			// during the last phase
			ex.op = ADD
			ex.item = sign
			q.Pop()

		case MINUS_EXPR:
			if ex.op != BOOLEAN {
				return nil, fmt.Errorf("duplicate operation sign: %s and %s", MINUS_EXPR, ex.item)
			}

			ex.op = MINUS
			ex.item = sign
			q.Pop()

		case MULTI_EXPR:
			if ex.op != BOOLEAN {
				return nil, fmt.Errorf("duplicate operation sign: %s and %s", MULTI_EXPR, ex.item)
			}

			ex.op = MULTIPLY
			ex.item = sign
			q.Pop()

		case DIV_EXPR:
			if ex.op != BOOLEAN {
				return nil, fmt.Errorf("duplicate operation sign: %s and %s", DIV_EXPR, ex.item)
			}

			ex.op = DIVIDE
			ex.item = sign
			q.Pop()
		}
	}
	return nil, fmt.Errorf("error when trying to parse token %s, followed by: %s", sign, err)
}

// IS_VALID_NUMBER takes in a string and ch-
// eck if it is a valid number, i.e. there
// are no non-numerical letters with in and
// '.' only appear once
func IS_VALID_NUMBER(s string) bool {
	return false
}
