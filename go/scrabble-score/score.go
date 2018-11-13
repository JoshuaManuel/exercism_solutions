//Package scrabble implements a letter-scorer
package scrabble

import (
	"strings"
)

//A map of the scores for each letter
var letterScores = map[string]int{
	"AEIOULNRST": 1,
	"DG":         2,
	"BCMP":       3,
	"FHVWY":      4,
	"K":          5,
	"JX":         8,
	"QZ":         10}

//Score returns the scrabble score of the input
func Score(str string) int {

	sum := 0
	str = strings.ToUpper(str)
	for _, l := range str {
		for key, val := range letterScores {
			if strings.Contains(key, string(l)) {
				sum += val
				break
			}
		}
	}
	return sum
}
