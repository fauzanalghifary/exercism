package acronym

import (
	"regexp"
	"strings"
	"unicode"
)

func Abbreviate(s string) string {
	splitter := regexp.MustCompile(`\s|-`)
	arrayWord := splitter.Split(s, -1)
	result := ""

	for _, word := range arrayWord {
		if len(word) == 0 {
			continue
		}

		for _, char := range word {
			if unicode.IsLetter(char) {
				result += strings.ToUpper(string(char))
				break
			}
		}

	}

	return result
}
