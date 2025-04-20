package isogram

import "strings"

func IsIsogram(word string) bool {
	var (
		seen      = make(map[rune]bool)
		lowerWord = strings.ToLower(word)
	)

	for _, char := range lowerWord {
		if char == ' ' || char == '-' {
			continue
		}
		if seen[char] {
			return false
		}
		seen[char] = true
	}

	return true
}
