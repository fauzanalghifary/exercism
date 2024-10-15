package pascal

func Triangle(n int) [][]int {
	result := make([][]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < i+1; j++ {
			if i < 2 || j == 0 || j == i {
				result[i] = append(result[i], 1)
			} else {
				prevRow := result[i-1]
				result[i] = append(result[i], prevRow[j-1]+prevRow[j])
			}
		}
	}

	return result
}
