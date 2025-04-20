package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("DNA is not the same length")
	}

	var differences int

	for i := range a {
		if a[i] != b[i] {
			differences++
		}
	}

	return differences, nil
}
