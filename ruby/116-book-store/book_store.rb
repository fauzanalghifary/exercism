module BookStore
  BOOK_PRICE = 8.0
  DISCOUNTS = {
    1 => 0.0,
    2 => 0.05,
    3 => 0.10,
    4 => 0.20,
    5 => 0.25
  }

  def self.calculate_price(basket)
    return 0.0 if basket.empty?

    counts = Hash.new(0)
    basket.each { |book| counts[book] += 1 }

    frequencies = counts.values.sort.reverse
    min_price(frequencies)
  end

  private

  def self.min_price(frequencies)
    return 0.0 if frequencies.empty? || frequencies.all?(&:zero?)

    best_price = Float::INFINITY
    (1..5).each do |group_size|
      next if frequencies.length < group_size

      new_frequencies = frequencies.dup
      group_size.times { |i| new_frequencies[i] -= 1 }
      new_frequencies = new_frequencies.select { |f| f > 0 }.sort.reverse

      group_price = group_size * BOOK_PRICE * (1 - DISCOUNTS[group_size])
      total_price = group_price + min_price(new_frequencies)
      best_price = [best_price, total_price].min
    end

    best_price
  end
end
