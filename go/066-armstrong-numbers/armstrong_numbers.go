package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	stringNum := strconv.Itoa(n)
	power := float64(len(stringNum))
	var result float64

	for _, num := range stringNum {
		digit := float64(num - '0')
		result += math.Pow(digit, power)
	}

	return n == int(result)
}
