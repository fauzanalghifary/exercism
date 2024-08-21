package romannumerals

import "fmt"

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input > 3999 {
		return "", fmt.Errorf("input must be between 1 and 3999")
	}

	romanValues := []struct {
		value int
		char  string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var result string
	for _, rv := range romanValues {
		for input >= rv.value {
			result += rv.char
			input -= rv.value
		}
	}

	return result, nil
}
