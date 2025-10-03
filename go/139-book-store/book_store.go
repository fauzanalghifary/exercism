package bookstore

import "sort"

func Cost(books []int) int {
	if len(books) == 0 {
		return 0
	}

	// Count frequency of each book
	freq := make(map[int]int)
	for _, book := range books {
		freq[book]++
	}

	// Convert frequency map to slice of counts and sort for better grouping
	counts := make([]int, 0, len(freq))
	for _, count := range freq {
		counts = append(counts, count)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))

	// Use dynamic programming approach
	return minCostDP(counts)
}

func minCostDP(counts []int) int {
	// Remove zeros
	filtered := []int{}
	for _, c := range counts {
		if c > 0 {
			filtered = append(filtered, c)
		}
	}
	counts = filtered

	if len(counts) == 0 {
		return 0
	}

	// Try to balance groups optimally
	// The key insight: groups of 4 are better than groups of 5 and 3
	return calculateOptimal(counts)
}

func calculateOptimal(counts []int) int {
	if len(counts) == 0 {
		return 0
	}

	total := 0
	groups := make([]int, 6) // groups[i] = number of groups of size i

	for len(counts) > 0 {
		// Count how many different books we have with at least 1 copy
		distinctBooks := len(counts)

		if distinctBooks >= 5 {
			groups[5]++
			for i := 0; i < 5; i++ {
				counts[i]--
			}
		} else if distinctBooks == 4 {
			groups[4]++
			for i := 0; i < 4; i++ {
				counts[i]--
			}
		} else if distinctBooks == 3 {
			groups[3]++
			for i := 0; i < 3; i++ {
				counts[i]--
			}
		} else if distinctBooks == 2 {
			groups[2]++
			for i := 0; i < 2; i++ {
				counts[i]--
			}
		} else {
			groups[1]++
			counts[0]--
		}

		// Remove zeros
		filtered := []int{}
		for _, c := range counts {
			if c > 0 {
				filtered = append(filtered, c)
			}
		}
		counts = filtered
	}

	// Optimize: convert pairs of (5, 3) to pairs of (4, 4)
	// because 2*4*640 = 5120 < 5*750 + 3*720 = 5160
	swaps := min(groups[5], groups[3])
	groups[5] -= swaps
	groups[3] -= swaps
	groups[4] += 2 * swaps

	// Calculate total cost
	for size := 1; size <= 5; size++ {
		total += groups[size] * groupPrice(size)
	}

	return total
}

func groupPrice(size int) int {
	basePrice := 800
	switch size {
	case 1:
		return basePrice
	case 2:
		return 2 * basePrice * 95 / 100
	case 3:
		return 3 * basePrice * 90 / 100
	case 4:
		return 4 * basePrice * 80 / 100
	case 5:
		return 5 * basePrice * 75 / 100
	default:
		return 0
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
