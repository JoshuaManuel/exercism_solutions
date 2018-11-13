//Package hamming implements a function to do check hamming distance
package hamming

import (
	"errors"
)

/*
Distance returns the Hamming distance of two strings
*/
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("the two strings were not the same length")
	}
	var errorSum int
	for c, l := range a {
		if l != rune(b[c]) {
			errorSum++
		}
	}
	return errorSum, nil
}
