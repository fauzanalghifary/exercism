class BinarySearch
  attr_reader :input

  # @param [Array] input
  def initialize(input)
    @input = input
  end

  # @param [Numeric] num
  def search_for(num, start_index = 0, end_index = input.size - 1)
    middle_index = ((end_index + start_index) / 2).floor

    return nil if middle_index < start_index || middle_index > end_index

    if num == input[middle_index]
      middle_index
    elsif num > input[middle_index]
      search_for(num, middle_index + 1, end_index)
    elsif num < input[middle_index]
      search_for(num, start_index, middle_index - 1)
    end
  end
end