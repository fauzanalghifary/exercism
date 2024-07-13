class Transpose
  def self.transpose(input)
    return '' if input == ''

    input_array = input.split("\n")
    max_length = input_array[0].size
    input_array.each do |input|
      max_length = input.size if input.size > max_length
    end

    result = []
    (0...max_length).each do |index|
      temp = ''
      last_index = 0
      input_array.each_with_index do |input, idx|
        if input[index]
          temp += input[index]
          last_index = idx
        else
          temp += ' '
        end
      end

      result.push(temp.slice(0, last_index + 1))
    end
    result.join("\n")
  end


  # def self.transpose(input)
  #   strings = input.gsub(' ', '_').split("\n").map(&:chars)
  #   max_length = strings.map(&:length).max
  #   strings.map { |line| line + ([' '] * (max_length - line.length)) }
  #          .transpose
  #          .map { |line| line.join.rstrip.gsub('_', ' ') }
  #          .join("\n")
  # end
end