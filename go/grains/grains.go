//Package grains implements two functions to count the grains on a chessboard
package grains

import (
	"errors"
)

//Square computes the number of grains that will be delivered on a certain day
func Square(num int) (uint64, error) {
	if num < 1 || num > 64 {
		return 0, errors.New("number was out of bounds (valid numbers are between 1 and 64)")
	}
	return 1 << uint64(num-1), nil
}

//Total will compute the total number of grains by the 64th day
func Total() uint64 {
	return 1<<uint64(64) - 1
}
