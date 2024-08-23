package isbn

import (
	"regexp"
	"strings"
)

func IsValidISBN(isbn string) bool {
	strippedInput := strings.ReplaceAll(isbn, "-", "")
	validISBN := regexp.MustCompile(`^\d-?\d{3}-?\d{5}-?\d$|X$`)

	if !validISBN.MatchString(isbn) {
		return false
	}

	var sum int
	for i, char := range strippedInput {
		if string(char) == "X" {
			sum += 10
		} else {
			sum += int(char-'0') * (10 - i)
		}
	}
	return sum%11 == 0
}
