package collatzconjecture

import "errors"

// Take any positive integer n.
// If n is even, divide n by 2 to get n / 2.
// If n is odd, multiply n by 3 and add 1 to get 3n + 1.
// Repeat the process indefinitely.
// The conjecture states that no matter which number you start with,
// you will always reach 1 eventually.
func CollatzConjecture(n int) (int, error) {
	var counter = 0
	if n <= 0 {
		return 0, errors.New("value is not positive")
	}

	for n != 1 {
		counter++
		if n%2 == 0 {
			n = n / 2
			continue
		}
		n = n*3 + 1
	}
	return counter, nil
}
