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
