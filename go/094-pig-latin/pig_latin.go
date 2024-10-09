package piglatin

import (
	"slices"
	"strings"
)

var vowels = []string{"a", "e", "i", "o", "u"}

func Sentence(sentence string) string {
	words := strings.Split(sentence, " ")
	var output []string

	for _, word := range words {
		output = append(output, ConvertWord(word))
	}

	return strings.Join(output, " ")
}

func ConvertWord(sentence string) string {
	firstChar := string(sentence[0])
	firstTwoChar := firstChar + string(sentence[1])
	if isVowel(firstChar) || firstTwoChar == "xr" || firstTwoChar == "yt" {
		return sentence + "ay"
	}

	var consonantStack string
	for i, char := range sentence {
		if isVowel(string(char)) || (char == 'y' && i != 0) {
			return sentence[i:] + consonantStack + "ay"
		}

		consonantStack += string(char)

		if string(char) == "q" && i != len(sentence)-1 && sentence[i+1] == 'u' {
			return sentence[i+2:] + sentence[0:i+2] + "ay"
		}
	}

	return sentence
}

func isVowel(char string) bool {
	if slices.Contains(vowels, char) {
		return true
	}

	return false
}
