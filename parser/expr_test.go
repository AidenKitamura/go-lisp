package parser

import (
	"fmt"
	"testing"
)

func TestExprADD2(t *testing.T) {
	op1 := &EXPR{INT, "2", nil}
	op2 := &EXPR{INT, "3", nil}
	op3 := EXPR{ADD, "+", []*EXPR{op1, op2}}
	tp, val := EVAL(&op3)
	if tp == ERROR {
		t.Fatalf("found type to be error. Value is: %v", val)
	}
	if val.(int) != 5 {
		t.Fatalf("found value error. Value expected to be 5, got: %v", val)
	}
	fmt.Printf("correct: %d %v\n", tp, val)
}

func TestExprADD4(t *testing.T) {
	op1 := &EXPR{INT, "2", nil}
	op2 := &EXPR{INT, "3", nil}
	op3 := &EXPR{ADD, "+", []*EXPR{op1, op2}}
	op4 := &EXPR{ADD, "+", []*EXPR{op1, op2}}
	op5 := &EXPR{ADD, "+", []*EXPR{op3, op4}}
	op6 := &EXPR{ADD, "+", []*EXPR{op5}}
	tp, val := EVAL(op6)
	if tp == ERROR {
		t.Fatalf("found type to be error. Value is: %v", val)
	}
	if val.(int) != 10 {
		t.Fatalf("found value error. Value expected to be 5, got: %v", val)
	}
	fmt.Printf("correct: %d %v\n", tp, val)
}

func TestExprADDMINUS4(t *testing.T) {
	op0 := &EXPR{FLOAT, "123.98", nil}
	op1 := &EXPR{INT, "2", nil}
	op2 := &EXPR{FLOAT, "3.765", nil}
	op3 := &EXPR{ADD, "+", []*EXPR{op1, op2}}
	op4 := &EXPR{ADD, "+", []*EXPR{op1, op2}}
	op5 := &EXPR{MINUS, "-", []*EXPR{op3, op4, op0}}
	tp, val := EVAL(op5)
	if tp == ERROR {
		t.Fatalf("found type to be error. Value is: %v", val)
	}
	if val.(float64) != (2+3.765)-2-3.765-123.98 {
		t.Fatalf("Invalid value %v", val)
	}
	fmt.Printf("correct: %d %v\n", tp, val)
}

func TestExprMULTI2(t *testing.T) {
	op0 := &EXPR{FLOAT, "123.98", nil}
	op1 := &EXPR{INT, "2", nil}
	op2 := &EXPR{MULTIPLY, "*", []*EXPR{op1, op0}}
	tp, val := EVAL(op2)
	if tp == ERROR {
		t.Fatalf("found type to be error. Value is: %v", val)
	}
	if val.(float64) != 123.98*2 {
		t.Fatalf("Invalid value %v", val)
	}
}

func TestExprDIVIDE2(t *testing.T) {
	op0 := &EXPR{FLOAT, "123.98", nil}
	op1 := &EXPR{INT, "2", nil}
	op2 := &EXPR{DIVIDE, "/", []*EXPR{op0, op1}}
	tp, val := EVAL(op2)
	if tp == ERROR {
		t.Fatalf("found type to be error. Value is: %v", val)
	}
	if val.(float64) != 123.98/2 {
		t.Fatalf("Invalid value %v", val)
	}
}
