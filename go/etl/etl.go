//Package etl the solution the etl exercism problem
package etl

import "strings"

//Transform implements takes the input map and transforms it to the correct format
func Transform(input map[int][]string) (output map[string]int) {
	output = make(map[string]int)
	for key, val := range input {
		for _, letter := range val {
			output[strings.ToLower(letter)] = key
		}
	}
	return output
}
