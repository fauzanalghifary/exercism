package railfence

func Encode(message string, rails int) string {
	hash := make(map[int]string)
	for i := 0; i < rails; i++ {
		hash[i+1] = ""
	}

	var index = 1
	var isIncreasing = true
	for _, char := range message {
		hash[index] += string(char)

		if index == 1 {
			index++
			isIncreasing = true
		} else if index == rails {
			index--
			isIncreasing = false
		} else if isIncreasing {
			index++
		} else {
			index--
		}
	}

	var result string
	for i := 0; i < rails; i++ {
		result += hash[i+1]
	}

	return result
}

func Decode(message string, rails int) string {
	divisor := (rails-2)*2 + 2
	chunkLength := len(message) / divisor
	remainder := len(message) % divisor

	hash := make(map[int]string)
	index := 0
	for i := 0; i < rails; i++ {
		endPoint := index + chunkLength
		if i != 0 && i+1 != rails {
			endPoint = index + (chunkLength * 2)
		}

		if remainder > 0 {
			endPoint++
			remainder--
		}

		hash[i+1] = message[index:endPoint]
		index = endPoint
	}

	var result string
	idx := 1
	isIncreasing := true
	for len(result) < len(message) {
		result += string(hash[idx][0])
		hash[idx] = hash[idx][1:]

		if idx == 1 {
			isIncreasing = true
		} else if idx == rails {
			isIncreasing = false
		}

		if isIncreasing {
			idx++
		} else {
			idx--
		}
	}

	return result
}

/*

2 rails -> (2-2)2+2 = 2
3 rails -> (3-2)2+2 = 4
4 rails -> (4-2)2+2 = 6
5 rails -> (5-2)2+2 = 8
6 rails -> (6-2)2+2 = 10

25 / 4 = 6
rails 1 = 6
rails 2 = 6+6
rails 3 = 6

5 rails => 8
17 / 8 = 2
17 % 8 = 1

23 / 8 = 2
23 % 8 = 7


*/
