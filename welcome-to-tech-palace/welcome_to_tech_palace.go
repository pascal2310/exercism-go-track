package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return fmt.Sprintf("Welcome to the Tech Palace, %s", strings.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	border := ""
	for _ = range make([]string, numStarsPerLine) {
		border += "*"
	}
	return fmt.Sprintf("%s\n%s\n%s", border, welcomeMsg, border)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	withoutStars := strings.ReplaceAll(oldMsg, "*", "")
	withoutNewLines := strings.ReplaceAll(withoutStars, "\n", "")
	return strings.TrimSpace(withoutNewLines)
}
