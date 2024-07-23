class Sieve
  attr_reader :num
  def initialize(num)
    @num = num
  end

  def primes
    hash_num = (2..num).to_a.to_h { |item| [item, false] }
    hash_num.each_with_object([]) do |(n, marked), primes_num|
      next if marked

      primes_num.push(n)

      multiplier = 2
      while n * multiplier <= num
        hash_num[n * multiplier] = true
        multiplier += 1
      end
    end
  end
end