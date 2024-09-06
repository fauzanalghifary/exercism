package sieve

func Sieve(limit int) []int {
	var prime []int
	markedNum := make(map[int]bool)

	for i := 2; i <= limit; i++ {
		if markedNum[i] == true {
			continue
		}

		prime = append(prime, i)
		multiplier := 2

		for i*multiplier <= limit {
			markedNum[i*multiplier] = true
			multiplier += 1
		}
	}

	return prime
}
