package change

import "errors"

func Change(coins []int, target int) ([]int, error) {
	if target < 0 {
		return nil, errors.New("negative target")
	}

	if target == 0 {
		return []int{}, nil
	}

	// dp[i] represents the minimum number of coins needed to make amount i
	dp := make([]int, target+1)
	// parent[i] stores which coin was used to reach amount i
	parent := make([]int, target+1)

	// Initialize dp array with impossible values
	for i := 1; i <= target; i++ {
		dp[i] = target + 1
		parent[i] = -1
	}

	// Build up solutions for each amount from 1 to target
	for i := 1; i <= target; i++ {
		for _, coin := range coins {
			if coin <= i && dp[i-coin]+1 < dp[i] {
				dp[i] = dp[i-coin] + 1
				parent[i] = coin
			}
		}
	}

	// If dp[target] is still impossible, no solution exists
	if dp[target] > target {
		return nil, errors.New("no combination can add up to target")
	}

	// Reconstruct the solution by backtracking through the parent array
	result := []int{}
	amount := target
	for amount > 0 {
		coin := parent[amount]
		result = append(result, coin)
		amount -= coin
	}

	// Sort the result in ascending order
	for i := 0; i < len(result)-1; i++ {
		for j := i + 1; j < len(result); j++ {
			if result[i] > result[j] {
				result[i], result[j] = result[j], result[i]
			}
		}
	}

	return result, nil
}
