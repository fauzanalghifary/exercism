package atbash

import (
	"unicode"
)

func Atbash(s string) string {
	var result string
	for i, char := range s {
		if unicode.IsLetter(char) {
			result += string(219 - unicode.ToLower(char))
		} else if unicode.IsDigit(char) {
			result += string(unicode.ToLower(char))
		} else {
			continue
		}

		if len(result)%6 == 5 && i < len(s)-2 {
			result += " "
		}

	}

	return result
}

/*
97 -> 122
98 -> 121
...
121 -> 98
122 -> 97
*/
