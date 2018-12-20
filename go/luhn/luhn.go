//Package luhn implements the function to check if a credit card number is valid
package luhn

import (
	"unicode"
)

//Valid checks whether the input number is a valid credit card number according to the luhn algorithm. A valid number will only have spaces and digits.
func Valid(input string) bool {
	shouldDouble := false
	sum := 0
	digits := 0 //choose to count loops because the input may have spaces or other non-digits
	for i := len(input) - 1; i >= 0; i-- {

		char := rune(input[i])

		if unicode.IsSpace(char) {
			continue
		}

		//if the character isn't a digit or space, then the input is malformed!
		if !unicode.IsDigit(char) {
			return false
		}

		//get the value of what can now only be a digit
		val := input[i] - '0'

		if shouldDouble {
			val *= 2
			if val > 9 {
				val = val - 9
			}
		}

		sum += int(val)

		shouldDouble = !shouldDouble
		digits++
	}

	if digits < 2 {
		return false
	}
	if sum%10 == 0 {
		return true
	}
	return false
}
