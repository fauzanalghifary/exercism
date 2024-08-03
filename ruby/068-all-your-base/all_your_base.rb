class BaseConverter
  # @param [Numeric] input_base
  # @param [Array<Numeric>] digits
  # @param [Numeric] output_base
  # @return [Array<Numeric>]
  def self.convert(input_base, digits, output_base)
    raise ArgumentError if input_base < 2 || output_base < 2

    base_10_num = 0
    digits.reverse.each_with_index do |num, index|
      raise ArgumentError if num.negative? || num >= input_base

      base_10_num += num * input_base**index
    end

    return [0] if base_10_num.zero?

    max_digit = log(output_base, base_10_num).floor
    result = []

    (0..max_digit).reverse_each do |num|
      subtract = (base_10_num / output_base**num).floor
      result << subtract
      base_10_num -= (subtract * output_base**num)
    end

    result
  end

  def self.log(base, number)
    Math.log(number) / Math.log(base)
  end
end