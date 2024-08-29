package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	delimiter := regexp.MustCompile(`[^a-zA-Z0-9']`)
	wordArray := delimiter.Split(phrase, -1)
	freq := make(Frequency)

	for _, word := range wordArray {
		word := strings.Trim(word, "'")
		if len(word) == 0 {
			continue
		}

		freq[strings.ToLower(word)]++
	}

	return freq
}

// Community Solution
//func WordCount(sentence string) Frequency {
//	result := make(Frequency)
//	reg := regexp.MustCompile(`\w+('\w+)?`)
//	for _, word := range reg.FindAllString(strings.ToLower(sentence), -1) {
//		result[word]++
//	}
//	return result
//}
