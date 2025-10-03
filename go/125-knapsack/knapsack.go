package knapsack

type Item struct {
	Weight, Value int
}

// Knapsack takes in a maximum carrying capacity and a collection of items
// and returns the maximum value that can be carried by the knapsack
// given that the knapsack can only carry a maximum weight given by maximumWeight
func Knapsack(maximumWeight int, items []Item) int {
	n := len(items)
	if n == 0 || maximumWeight == 0 {
		return 0
	}

	// Create a 2D DP table where dp[i][w] represents the maximum value
	// that can be obtained with the first i items and weight limit w
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, maximumWeight+1)
	}

	// Fill the DP table
	for i := 1; i <= n; i++ {
		item := items[i-1]
		for w := 0; w <= maximumWeight; w++ {
			// Option 1: Don't take the current item
			dp[i][w] = dp[i-1][w]

			// Option 2: Take the current item (if it fits)
			if item.Weight <= w {
				valueWithItem := dp[i-1][w-item.Weight] + item.Value
				if valueWithItem > dp[i][w] {
					dp[i][w] = valueWithItem
				}
			}
		}
	}

	return dp[n][maximumWeight]
}
