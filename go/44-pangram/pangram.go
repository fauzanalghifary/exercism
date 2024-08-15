package pangram

import (
	"strings"
)

func IsPangram(input string) bool {
	lowercaseInput := strings.ToLower(input)
	for char := 'a'; char <= 'z'; char++ {
		if !strings.ContainsRune(lowercaseInput, char) {
			return false
		}
	}
	return true
}
