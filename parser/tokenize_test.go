package parser

import (
	"fmt"
	"testing"
)

func TestTokenize(t *testing.T) {
	testList := []struct {
		in       string
		expected []string
	}{
		{"((((", []string{"(", "(", "(", "("}},
		{"(+12 23)((1 1))", []string{"(", "+", "12", "23", ")", "(", "(", "1", "1", ")", ")"}},
		{"", []string{}},
		{"1$@(()!@3", []string{"1$@", "(", "(", ")", "!@3"}},
		{"       1       2       3()", []string{"1", "2", "3", "(", ")"}},
		{"((\n))", []string{"(", "(", ")", ")"}},
	}

	for _, testCase := range testList {
		actual := testCase.expected
		have := Tokenize(testCase.in)
		fmt.Println(have)
		if len(actual) != len(have) {
			t.Fatal("different length of tokens", len(actual), len(have))
		}
		for i := 0; i < len(actual); i++ {
			if actual[i] != have[i] {
				t.Fatalf("%d-th token different: %s, %s", i, actual[i], have[i])
			}
		}
	}
}
