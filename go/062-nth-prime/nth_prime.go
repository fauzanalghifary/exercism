package prime

import (
	"errors"
	"math"
)

// Nth returns the nth prime number.
// An error must be returned if the nth prime number can't be calculated
// ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("invalid input")
	}

	upperBound := int(float64(n)*math.Log(float64(n)) + float64(n)*math.Log(math.Log(float64(n))))

	hashNum := make(map[int]bool)
	var primeNum []int
	for i := 2; len(primeNum) < n; i++ {
		if hashNum[i] == true {
			continue
		}

		primeNum = append(primeNum, i)

		multiplier := 2
		for i*multiplier <= upperBound {
			hashNum[i*multiplier] = true
			multiplier += 1
		}
	}

	return primeNum[n-1], nil
}
