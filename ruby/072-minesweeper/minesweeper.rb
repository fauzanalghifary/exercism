class Minesweeper
  # @param [Array<String>] input
  def self.annotate(input)
    result = Array.new(input.length) { Array.new(input[0].length, 0) }
    input.each_with_index do |row, row_index|
      row.each_char.with_index do |char, str_index|
        if char == "*"
          result[row_index - 1][str_index] += 1 if row_index > 0
          result[row_index + 1][str_index] += 1 if row_index < input.length - 1
          result[row_index][str_index - 1] += 1 if str_index > 0
          result[row_index][str_index + 1] += 1 if str_index < row.length - 1
          result[row_index - 1][str_index - 1] += 1 if row_index > 0 && str_index > 0
          result[row_index - 1][str_index + 1] += 1 if row_index > 0 && str_index < row.length - 1
          result[row_index + 1][str_index - 1] += 1 if row_index < input.length - 1 && str_index > 0
          result[row_index + 1][str_index + 1] += 1 if row_index < input.length - 1 && str_index < row.length - 1
        end
      end
    end

    input.each_with_index do |row, row_index|
      row.each_char.with_index do |char, str_index|
        if char == "*"
          result[row_index][str_index] = "*"
        end

        if result[row_index][str_index] == 0
          result[row_index][str_index] = " "
        end
      end
    end

    result.map! { |row| row.join }
  end
end
