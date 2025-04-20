package diffsquares

func SquareOfSum(n int) int {
	var sum int
	for i := range make([]string, n) {
		sum += i + 1
	}

	return sum * sum
}

func SumOfSquares(n int) int {
	var sum int
	for i := range make([]string, n+1) {
		sum += i * i
	}

	return sum
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
