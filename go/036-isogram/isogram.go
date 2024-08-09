package isogram

import (
	"unicode"
)

func IsIsogram(word string) bool {
	dict := map[rune]int{}
	//re := regexp.MustCompile(`\w`)

	for _, char := range word {
		//if !re.MatchString(string(char)) {
		if !unicode.IsLetter(char) {
			continue
		}

		key := unicode.ToUpper(char)
		if dict[key] != 0 {
			return false
		} else {
			dict[key] = 1
		}
	}

	return true
}
