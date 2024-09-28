package allyourbase

import (
	"errors"
	"math"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return []int(nil), errors.New("input base must be >= 2")
	}

	if outputBase < 2 {
		return []int(nil), errors.New("output base must be >= 2")
	}

	if len(inputDigits) == 0 || len(inputDigits) == 1 && inputDigits[0] == 0 {
		return []int{0}, nil
	}

	var sum int
	var power float64 = 0
	for i := len(inputDigits) - 1; i >= 0; i-- {
		if inputDigits[i] < 0 || inputDigits[i] >= inputBase {
			return []int(nil), errors.New("all digits must satisfy 0 <= d < input base")
		}

		sum += inputDigits[i] * int(math.Pow(float64(inputBase), power))
		power += 1
	}

	if sum == 0 {
		return []int{0}, nil
	}

	maxBase := int(math.Floor(math.Log(float64(sum)) / math.Log(float64(outputBase))))
	var remainder int = sum

	var output []int
	for i := maxBase; i >= 0; i-- {
		pow := int(math.Pow(float64(outputBase), float64(i)))
		digit := remainder / pow
		output = append(output, digit)
		remainder -= digit * pow
	}

	return output, nil
}
