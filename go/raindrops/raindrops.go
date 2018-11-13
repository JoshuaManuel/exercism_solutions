//Package raindrops implements the solution to the raindrops exercism problem
package raindrops

import (
	"strconv"
)

//Convert returns a silly string based on the factors of the number
func Convert(num int) string {
	var str string //declare the initial string

	if num%3 == 0 {
		str += "Pling"
	}

	if num%5 == 0 {
		str += "Plang"
	}

	if num%7 == 0 {
		str += "Plong"
	}

	if len(str) == 0 {
		return strconv.Itoa(num) //return the input as a string
	}

	return str
}
