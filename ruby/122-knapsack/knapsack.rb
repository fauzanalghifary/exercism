class Knapsack
  def initialize(max_weight)
    @max_weight = max_weight
  end

  def max_value(items)
    return 0 if items.empty? || @max_weight <= 0

    n = items.length
    # Create DP table: dp[i][w] = max value using first i items with weight limit w
    dp = Array.new(n + 1) { Array.new(@max_weight + 1, 0) }

    # Fill the DP table
    (1..n).each do |i|
      item = items[i - 1]
      (0..@max_weight).each do |w|
        # Don't take the item
        dp[i][w] = dp[i - 1][w]

        # Take the item if it fits
        if item.weight <= w
          value_with_item = dp[i - 1][w - item.weight] + item.value
          dp[i][w] = [dp[i][w], value_with_item].max
        end
      end
    end

    dp[n][@max_weight]
  end
end
