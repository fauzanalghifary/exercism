package luhn

import (
	"regexp"
	"strings"
)

func Valid(id string) bool {
	strippedWord := strings.ReplaceAll(id, " ", "")
	re := regexp.MustCompile(`\D`)

	if re.MatchString(strippedWord) || len(strippedWord) <= 1 {
		return false
	}

	sum := 0

	for i, char := range strippedWord {
		digit := int(char - '0')

		if (len(strippedWord)-i)%2 == 0 {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
	}

	return sum%10 == 0
}
