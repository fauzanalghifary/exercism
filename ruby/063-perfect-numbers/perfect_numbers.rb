class PerfectNumber
  # @param [Numeric] num
  def self.classify(num)
    raise RuntimeError if num.negative?

    factors = (2..Math.sqrt(num).to_i).each_with_object([1]) do |n, arr|
      if (num % n).zero?
        arr << n
        arr << num / n
      end
    end

    factors_sum = factors.sum

    return 'perfect' if factors_sum == num
    return 'abundant' if factors_sum > num

    'deficient'
  end
end