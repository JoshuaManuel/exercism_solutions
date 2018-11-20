//Package diffsquares implements some functions for calculating the difference of squares
package diffsquares

//SquareOfSum calculates sum of all the numbers from 0-num, then squares it
func SquareOfSum(num int) (sum int) {
	sum = num * (num + 1) / 2 //formula to calculate the sum on numbers 0-num
	return sum * sum
}

//SumOfSquares calculates the sum of all the squares from 0-n
func SumOfSquares(num int) (sum int) {
	return (num * (num + 1) * (2*num + 1)) / 6 //formula to calculate the sum of squares 0-num
}

//Difference calculates the difference between SquareOfSum and SumOfSquares when they are fed the same input
func Difference(num int) int {
	return SquareOfSum(num) - SumOfSquares(num)
}
