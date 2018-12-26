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
	var ans uint64 = 1
	for i := 0; i < num-1; i++ {
		ans = ans * 2
	}
	return ans, nil
}

//Total will compute the total number of grains by the 64th day
func Total() uint64 {
	var sum uint64
	for i := 1; i <= 64; i++ {
		squared, _ := Square(i)
		sum += squared
	}
	return sum
}
