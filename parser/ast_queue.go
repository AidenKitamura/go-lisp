// ast_queue provides a helper queue class
// to deal with ast generation. The ast q-
// ueue does no check on any possible err-
// ors, but just create the helper instan-
// ce. Pushing is not provided.

package parser

import "fmt"

type ASTQUEUE struct {
	ASTELEM []string
}

// Returns the first element currently in
// the queue, but do not eliminate it.
// If there is no such element, an empty
// string and error is returned.
func (A *ASTQUEUE) Top() (string, error) {
	if len(A.ASTELEM) == 0 {
		return "", fmt.Errorf("reading elements from an empty ast queue")
	}

	return A.ASTELEM[0], nil
}

// Returns the first element in the queue
// and delete it afterwards. If pop from
// an empty queue, an empty string is re-
// turned with an error.
func (A *ASTQUEUE) Pop() (string, error) {
	if len(A.ASTELEM) == 0 {
		return "", fmt.Errorf("popping elements from an empty ast queue")
	}

	topElem := A.ASTELEM[0]
	A.ASTELEM = A.ASTELEM[1:]
	return topElem, nil
}

// return the result if there is no more
// elements in the queue
func (A *ASTQUEUE) ReachedEnd() bool {
	return len(A.ASTELEM) == 0
}

// ASTQUEUE constructor using an existing
// token list, returns the pointer to the
// constructed astqueue
func ASTQUEUEConstructer(tokens []string) *ASTQUEUE {
	return &ASTQUEUE{tokens}
}
