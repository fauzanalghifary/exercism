class RunLengthEncoding
  # @param [String] input
  def self.encode(input)
    return '' if input.empty?

    result = ''
    current_char = input[0]
    current_count = 1

    input[1..].each_char do |char|
      if char == current_char
        current_count += 1
      else
        result += (current_count > 1 ? current_count.to_s : '') + current_char
        current_char = char
        current_count = 1
      end
    end

    result + (current_count > 1 ? current_count.to_s : '') + current_char
  end

  # @param [String] input
  def self.decode(input)
    result = ""
    current_char = ""

    input.each_char do |char|
      if self.numeric_string?(char)
        current_char += char
      else
        num = current_char.to_i.nonzero? || 1
        num.times {|_| result += char }
        current_char = ""
      end
    end

    result
  end

  private

  def self.numeric_string?(str)
    str.to_i.to_s == str
  end

  # def self.encode(input)
  #   input.gsub(/(.)\1+/) { |m| "#{m.length}#{m[0]}" }
  # end
  #
  # def self.decode(input)
  #   input.gsub(/\d+./) { |m| m[-1] * m.to_i }
  # end
end