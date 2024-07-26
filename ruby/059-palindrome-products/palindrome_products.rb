require 'set'

class Palindromes
  attr_reader :min_factor, :max_factor

  Palindrome = Struct.new(:value, :factors)

  def initialize(max_factor:, min_factor: 1)
    @max_factor = max_factor
    @min_factor = min_factor
  end

  # @return [Array<Hash>]
  def generate
    (min_factor..max_factor).each_with_object([]) do |num1, result|
      (min_factor..max_factor).each do |num2|
        product = num1 * num2
        next unless palindrome?(product)

        existing_product = result.find { |hash| hash[:value] == product }
        if existing_product
          existing_product[:factors] << [num1, num2]
        else
          the_hash = {
            value: product,
            factors: [[num1, num2]]
          }
          result << the_hash
        end
      end
    end
  end

  def largest
    result = generate.max_by { |item| item[:value] }
    filtered_factors = remove_duplicate_sum(result[:factors])
    Palindrome.new(value: result[:value], factors: filtered_factors)
  end

  def smallest
    result = generate.min_by { |item| item[:value] }
    filtered_factors = remove_duplicate_sum(result[:factors])
    Palindrome.new(value: result[:value], factors: filtered_factors)
  end

  private
  # @param [Integer] num
  def palindrome?(num)
    num.to_s == num.to_s.reverse
  end

  def remove_duplicate_sum(arr)
    seen_sums = Set.new
    arr.select do |array|
     sum = array.sum
     if seen_sums.include?(sum)
       false # Remove the sub-array
     else
       seen_sums.add(sum)
       true # Keep the sub-array
     end
   end
  end
end


# COMMUNITY SOLUTION

# class Palindromes
#   Palindrome = Struct.new(:value, :factors)
#   attr_accessor :largest, :smallest
#   def initialize(max_factor: 1, min_factor: 1)
#     @factors = [*min_factor..max_factor].repeated_combination(2)
#   end
#   def generate
#     @smallest, @largest = @factors
#                             .group_by { |arr| arr.reduce(:*) }
#                             .select { |x, _| palindrome? x }
#                             .minmax
#                             .map { |x, y| Palindrome.new(x, y) }
#   end
#   private
#   def palindrome?(x)
#     x.to_s.reverse == x.to_s
#   end
# end