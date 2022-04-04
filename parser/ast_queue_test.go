package parser

import (
	"strconv"
	"testing"
)

func TestAstQueue(t *testing.T) {
	testString := make([]string, 0)
	for i := 0; i < 20; i++ {
		testString = append(testString, strconv.Itoa(i))
	}
	testQueue := ASTQUEUEConstructer(testString)
	for i := 0; i < 20; i++ {
		testElem, err := testQueue.Pop()
		if err != nil || testElem != strconv.Itoa(i) {
			t.Fatal("incorrect elements: ", i, testElem)
		}
	}
	_, err := testQueue.Top()
	if err == nil {
		t.Fatal("expecting err, but getting nil")
	}
}
