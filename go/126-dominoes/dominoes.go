package dominoes

type Domino [2]int

func MakeChain(input []Domino) ([]Domino, bool) {
	if len(input) == 0 {
		return []Domino{}, true
	}

	if len(input) == 1 {
		if input[0][0] == input[0][1] {
			return input, true
		}
		return nil, false
	}

	// Try to build a chain using backtracking
	used := make([]bool, len(input))
	chain := []Domino{}

	for i := 0; i < len(input); i++ {
		// Try starting with domino in normal orientation
		used[i] = true
		chain = []Domino{input[i]}
		if backtrack(input, used, &chain) {
			return chain, true
		}

		// Try starting with domino flipped
		chain = []Domino{Domino{input[i][1], input[i][0]}}
		if backtrack(input, used, &chain) {
			return chain, true
		}
		used[i] = false
	}

	return nil, false
}

func backtrack(input []Domino, used []bool, chain *[]Domino) bool {
	if len(*chain) == len(input) {
		// Check if the chain forms a loop
		return (*chain)[0][0] == (*chain)[len(*chain)-1][1]
	}

	lastValue := (*chain)[len(*chain)-1][1]

	for i := 0; i < len(input); i++ {
		if used[i] {
			continue
		}

		// Try normal orientation
		if input[i][0] == lastValue {
			used[i] = true
			*chain = append(*chain, input[i])
			if backtrack(input, used, chain) {
				return true
			}
			*chain = (*chain)[:len(*chain)-1]
			used[i] = false
		}

		// Try flipped orientation
		if input[i][1] == lastValue {
			used[i] = true
			*chain = append(*chain, Domino{input[i][1], input[i][0]})
			if backtrack(input, used, chain) {
				return true
			}
			*chain = (*chain)[:len(*chain)-1]
			used[i] = false
		}
	}

	return false
}
