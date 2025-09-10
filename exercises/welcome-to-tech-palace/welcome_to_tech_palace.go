package techpalace

import (
	"fmt"
	"strings"
	"regexp"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return fmt.Sprintf("Welcome to the Tech Palace, %s", strings.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	stars := strings.Repeat("*", numStarsPerLine)
	return fmt.Sprintf("%s\n%s\n%s", stars, welcomeMsg, stars)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	start, err := regexp.Compile("^(\\*|\\s)+")
	if err != nil {
		panic(err)
	}

	end, err := regexp.Compile("(\\s|\\*)+$")
	if err != nil {
		panic(err)
	}

	newMsg := start.ReplaceAllString(oldMsg, "")
	newMsg = end.ReplaceAllString(newMsg, "")

	return newMsg
}
