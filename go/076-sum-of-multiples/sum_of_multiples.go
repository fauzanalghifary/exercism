package summultiples

func SumMultiples(limit int, divisors ...int) int {
	var sum int
	mark := make(map[int]bool)

	for _, num := range divisors {
		multiplier := 1
		for num*multiplier < limit && num > 0 {
			value := num * multiplier
			if mark[value] != true {
				mark[value] = true
				sum += value
			}
			multiplier += 1
		}
	}

	return sum
}
