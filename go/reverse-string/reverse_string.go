//Package reverse implements the solution to the reverse exercism problem
package reverse

//String returns a reversed version of the input string
func String(input string) string {
	chars := []rune(input)
	var reversed []rune
	for i := len(chars) - 1; i >= 0; i-- {
		reversed = append(reversed, chars[i])
	}
	return string(reversed)
}
