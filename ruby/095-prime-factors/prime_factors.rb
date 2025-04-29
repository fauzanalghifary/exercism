class PrimeFactors
  # @param[Numeric]num
  def self.of(num)
    return [] if num == 1

    factors = []
    remainder = num
    divisor = 2
    while remainder > 1
      if remainder % divisor == 0
        factors << divisor
        remainder /= divisor
      else
        divisor += 1
      end
    end

    factors
  end
end