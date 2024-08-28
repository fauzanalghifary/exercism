package anagram

import (
	"strings"
	"unicode"
)

func Detect(subject string, candidates []string) []string {
	var result []string

	for _, candidate := range candidates {
		if isAnagram(subject, candidate) {
			result = append(result, candidate)
		}
	}

	return result
}

func isAnagram(subject string, candidate string) bool {
	if strings.EqualFold(subject, candidate) {
		return false
	}

	if len(subject) != len(candidate) {
		return false
	}

	//var obj1 map[rune]int //CAN'T DO THIS
	//var obj2 map[rune]int //CAN'T DO THIS
	obj1 := make(map[rune]int)
	obj2 := make(map[rune]int)

	for _, char := range subject {
		obj1[unicode.ToLower(char)]++
	}

	for _, char := range candidate {
		obj2[unicode.ToLower(char)]++
	}

	for r, count := range obj1 {
		if obj2[r] != count {
			return false
		}
	}

	return true
}
