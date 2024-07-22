class PythagoreanTriplet
  # @param [String] n
  def self.triplets_with_sum(sum)
    (1..(sum / 3).floor).each_with_object([]) do |a, result|
      (a + 1..(sum / 2).floor).each do |b|
        c = sum - a - b
        result.push([a, b, c]) if a**2 + b**2 == c**2
      end
    end
  end
end

# class PythagoreanTriplet
#   def self.triplets_with_sum(sum)
#     (1..sum/3).each_with_object([]) do |a, triplets_set|
#       b = (sum * (sum - 2 * a)) / (2 * (sum - a))
#       c = sum - a - b
#       next if b < a || c < b
#       triplets_set << [a, b, c] if (a**2 + b**2 == c**2)
#     end
#   end
# end