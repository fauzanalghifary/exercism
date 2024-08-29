package lsproduct

import (
	"errors"
	"regexp"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if len(digits) < span || span < 1 {
		return 0, errors.New("invalid input")
	}

	containNotNumber := regexp.MustCompile(`\D`)
	if containNotNumber.MatchString(digits) {
		return 0, errors.New("invalid input")
	}

	var maxSum int
	var sum int
	for i := 0; i <= len(digits)-span; i++ {
		sum = 1
		for j := i; j < span+i; j++ {
			sum *= charToInt(digits[j])
		}

		if sum > maxSum {
			maxSum = sum
		}
	}

	return int64(maxSum), nil
}

func charToInt(char uint8) int {
	return int(char - '0')
}
