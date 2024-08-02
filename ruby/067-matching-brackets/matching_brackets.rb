class Brackets
  # @param [String] input
  def self.paired?(input)
    valid_char = %w([ ] { } ( ))
    closing_bracket = %w(] } \))
    next_chars = []

    input.each_char do |char|
      next unless valid_char.include?(char)

      next_char = next_chars.last

      if closing_bracket.include?(char)
        return false if char != next_char

        next_chars.pop
      end

      case char
      when '['
        next_chars << ']'
      when '{'
        next_chars << '}'
      when '('
        next_chars << ')'
      end

    end

    next_chars.empty?
  end
end