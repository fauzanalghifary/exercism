package scrabble

import (
	"strings"
	"unicode"
)

var letterScore = map[string]int{
	"AEIOULNRST": 1,
	"DG":         2,
	"BCMP":       3,
	"FHVWY":      4,
	"K":          5,
	"JX":         8,
	"QZ":         10,
}

func Score(word string) int {
	currentScore := 0
	for _, char := range word {
		for key, value := range letterScore {
			if strings.ContainsRune(key, unicode.ToUpper(char)) {
				currentScore += value
			}
		}
	}
	return currentScore
}
