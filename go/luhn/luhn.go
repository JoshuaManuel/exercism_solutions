//Package luhn implements the function to check if a credit card number is valid
package luhn

import (
	"unicode"
)

//Valid checks whether the input number is a valid credit card number according to the luhn algorithm. A valid number will only have spaces and digits.
func Valid(input string) bool {
	shouldDouble := false
	sum := 0
	loops := 0 //choose to count loops because the input may have spaces or other non-digits
	for i := len(input) - 1; i >= 0; i-- {

		//if the character isn't a digit or space, then the input is malformed!
		if !(unicode.IsDigit(rune(input[i])) || unicode.IsSpace(rune(input[i]))) {
			return false
		}

		//we can have spaces in the card number; just make sure we don't count it as a digit
		if unicode.IsSpace(rune(input[i])) {
			continue
		}
		//get the value of what can now only be a digit, not the
		val := input[i] - '0'

		if shouldDouble {
			/* this implementation is less apparent to the programmer
			if val != 9 {
				val = (val * 2) % 9
			}
			*/
			val = val * 2
			if val > 9 {
				val = val - 9
			}
		}

		sum += int(val)

		//toggle the shouldDouble variable
		if shouldDouble {
			shouldDouble = false
		} else {
			shouldDouble = true
		}
		loops++
	}

	if loops < 2 {
		return false
	}
	if sum%10 == 0 {
		return true
	}
	return false
}
