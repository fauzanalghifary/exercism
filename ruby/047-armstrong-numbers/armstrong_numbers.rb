module ArmstrongNumbers
  def self.include?(number)
    string_num = number.to_s
    num_of_digit = string_num.size
    string_num.chars.sum { |num| num.to_i**num_of_digit } == number

    #   number.digits.sum { |n| n ** number.digits.size } == number
  end
end
