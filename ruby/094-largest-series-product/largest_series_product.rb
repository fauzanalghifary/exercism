class Series
  # @param [String] num
  def initialize(num)
    raise ArgumentError, 'Digits input must only contain digits' unless num.match?(/^\d*$/)
    @digits = num.chars.map(&:to_i)
  end

  # @param [Integer] span
  # @return [Integer]
  def largest_product(span)
    raise ArgumentError, 'Span must be non-negative' if span.negative?
    raise ArgumentError, 'Span cannot be larger than digits length' if span > @digits.length

    @digits
      .each_cons(span)
      .map { |slice| slice.reduce(:*) }
      .max
  end
end


# class Series
#   # @param[String]num
#   def initialize(num)
#     raise ArgumentError if num.match /\D/
#     @num = num
#   end
#
#   # @param[Numeric]span
#   def largest_product(span)
#     raise ArgumentError if span > @num.length || span < 0
#     max = nil
#
#     @num.each_char.each_with_index do |c, index|
#       break if index == @num.length - span + 1
#       res = 1
#       @num[index..index+span-1].each_char do |n|
#         res *= n.to_i
#       end
#       max = res if max.nil? || res > max
#     end
#
#     max
#   end
# end