// tokenize includes functions to help create
// a list of tokens from a given string
package parser

// Tokenize is a function used to create the
// list of tokens to work with
func Tokenize(s string) []string {
	// Edge case 1
	// empty string
	if len(s) == 0 {
		return make([]string, 0)
	}

	// Other cases
	// We don't deal with validity
	// Only tokenize
	tokens := make([]string, 0)
	curPos := 0
	for pos, r := range s {
		switch r {
		case BEGIN_TOKEN, END_TOKEN, ADD_TOKEN, MINUS_TOKEN, MULTI_TOKEN, DIV_TOKEN:
			if curPos != pos {
				tokens = append(tokens, s[curPos:pos])
			}
			tokens = append(tokens, string(r))
			curPos = pos + 1
		case ' ', '\n':
			if curPos != pos {
				tokens = append(tokens, s[curPos:pos])
			}
			curPos = pos + 1
		}
	}
	// Right end situation
	if curPos != len(s) {
		tokens = append(tokens, s[curPos:])
	}
	return tokens
}
