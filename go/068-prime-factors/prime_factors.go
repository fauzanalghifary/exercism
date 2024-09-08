package prime

func Factors(n int64) []int64 {
	var factors []int64
	var divisor int64 = 2
	remaining := n

	for remaining > 1 {
		if remaining%divisor == 0 {
			factors = append(factors, divisor)
			remaining = remaining / divisor
		} else {
			divisor += 1
		}
	}

	return factors
}
