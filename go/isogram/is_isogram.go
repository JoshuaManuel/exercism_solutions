//Package isogram implements an isogram function
package isogram

import "unicode"

//IsIsogram returns whether or not the given string is an isogram
func IsIsogram(str string) bool {
	letters := make(map[rune]bool)
	for _, b := range str {
		b = unicode.ToUpper(b)
		if letters[b] && unicode.IsLetter(b) { //if it's a letter and already added
			return false
		}
		letters[b] = true
	}
	return true
}
