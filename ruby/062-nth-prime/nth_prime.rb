class Prime

  def self.nth(n)
    raise ArgumentError if n <= 0
    return [2, 3, 5, 7, 11, 13][n - 1] if n <= 6

    upper_bound = (n * Math.log(n) + n * Math.log(Math.log(n))).to_i

    sieve = Array.new(upper_bound + 1, true)
    sieve[0] = sieve[1] = false

    (2..Math.sqrt(upper_bound)).each do |i|
      next unless sieve[i]

      (i * i).step(upper_bound, i) do |j|
        sieve[j] = false
      end
    end

    primes = sieve.each_with_index.reduce([]) do |arr, (is_prime, i)|
      arr << i if is_prime
      arr
    end

    primes[n - 1]
  end


  # def self.nth n
  #   up_to = n * (Math.log(n) + 2)
  #   primes = (2..up_to).to_a
  #   primes.each {|num| primes.delete_if {|i| i > num && (i % num) == 0} }
  #   primes[n-1]
  # end
end