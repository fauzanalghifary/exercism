package yacht

func Score(dice []int, category string) int {
	var dicesMap = make(map[int]int)
	for _, d := range dice {
		dicesMap[d]++
	}

	switch category {
	case "ones":
		return dicesMap[1] * 1
	case "twos":
		return dicesMap[2] * 2
	case "threes":
		return dicesMap[3] * 3
	case "fours":
		return dicesMap[4] * 4
	case "fives":
		return dicesMap[5] * 5
	case "sixes":
		return dicesMap[6] * 6
	case "full house":
		if len(dicesMap) == 2 {
			sum := 0
			for diceNum, count := range dicesMap {
				if count != 2 && count != 3 {
					return 0
				}
				sum += diceNum * count
			}
			return sum
		}
		return 0
	case "four of a kind":
		for diceNum, count := range dicesMap {
			if count >= 4 {
				return diceNum * 4
			}
		}
		return 0
	case "little straight":
		if len(dicesMap) == 5 && dicesMap[6] == 0 {
			return 30
		}
		return 0
	case "big straight":
		if len(dicesMap) == 5 && dicesMap[1] == 0 {
			return 30
		}
		return 0
	case "choice":
		sum := 0
		for diceNum, count := range dicesMap {
			sum += diceNum * count
		}
		return sum
	case "yacht":
		if len(dicesMap) == 1 {
			return 50
		}
		return 0
	default:
		return 0
	}
}
