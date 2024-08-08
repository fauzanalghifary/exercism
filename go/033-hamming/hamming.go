package hamming

import (
	"errors"
	"unicode/utf8"
)

func Distance(a, b string) (int, error) {
	if utf8.RuneCountInString(a) != utf8.RuneCountInString(b) {
		return 0, errors.New("strands must be of equal length")
	}

	distance := 0
	strand1Runes := []rune(a)
	strand2Runes := []rune(b)

	for idx, rooney := range strand1Runes {
		if rooney != strand2Runes[idx] {
			distance++
		}
	}
	return distance, nil
}
