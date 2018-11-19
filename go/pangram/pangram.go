//Package pangram implements the solution to a Pangram checker
package pangram

import "unicode"

//IsPangram identifies if the given string is a pangram (uses every letter in the English alphabet)
func IsPangram(input string) bool {
	var letters = make(map[rune]bool)

	for _, val := range input {
		val = unicode.ToLower(val)
		if !letters[val] && unicode.IsLetter(val) {
			letters[val] = true
		}
	}

	if len(letters) == 26 { //if we have 26 letters in the letters map
		return true
	}
	return false
}
