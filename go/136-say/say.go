package say

import "strings"

func Say(n int64) (string, bool) {
	if n < 0 || n >= 1000000000000 {
		return "", false
	}

	if n == 0 {
		return "zero", true
	}

	return sayNumber(n), true
}

func sayNumber(n int64) string {
	var parts []string

	if n >= 1000000000 {
		billions := n / 1000000000
		parts = append(parts, sayHundreds(billions)+" billion")
		n %= 1000000000
	}

	if n >= 1000000 {
		millions := n / 1000000
		parts = append(parts, sayHundreds(millions)+" million")
		n %= 1000000
	}

	if n >= 1000 {
		thousands := n / 1000
		parts = append(parts, sayHundreds(thousands)+" thousand")
		n %= 1000
	}

	if n > 0 {
		parts = append(parts, sayHundreds(n))
	}

	return strings.Join(parts, " ")
}

func sayHundreds(n int64) string {
	var parts []string

	if n >= 100 {
		hundreds := n / 100
		parts = append(parts, ones[hundreds]+" hundred")
		n %= 100
	}

	if n >= 20 {
		tensDigit := n / 10
		onesDigit := n % 10
		if onesDigit > 0 {
			parts = append(parts, tensWords[tensDigit]+"-"+ones[onesDigit])
		} else {
			parts = append(parts, tensWords[tensDigit])
		}
		return strings.Join(parts, " ")
	}

	if n >= 10 && n < 20 {
		parts = append(parts, teens[n-10])
		return strings.Join(parts, " ")
	}

	if n > 0 {
		parts = append(parts, ones[n])
	}

	return strings.Join(parts, " ")
}

var ones = map[int64]string{
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
}

var teens = map[int64]string{
	0: "ten",
	1: "eleven",
	2: "twelve",
	3: "thirteen",
	4: "fourteen",
	5: "fifteen",
	6: "sixteen",
	7: "seventeen",
	8: "eighteen",
	9: "nineteen",
}

var tensWords = map[int64]string{
	2: "twenty",
	3: "thirty",
	4: "forty",
	5: "fifty",
	6: "sixty",
	7: "seventy",
	8: "eighty",
	9: "ninety",
}
