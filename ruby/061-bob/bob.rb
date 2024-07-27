class Bob
  # @param [String] input
  # @return [String]
  def self.hey(input)
    trimmed_input = input.gsub(/\s/, '')
    is_question = trimmed_input.match(/\?$/)
    is_yelling = trimmed_input == trimmed_input.upcase && trimmed_input.match(/[a-zA-Z]/)
    is_silence = !trimmed_input.match(/\w/)

    if is_yelling && is_question
      "Calm down, I know what I'm doing!"
    elsif is_yelling
      "Whoa, chill out!"
    elsif is_question
      "Sure."
    elsif is_silence
      "Fine. Be that way!"
    else
      "Whatever."
    end
  end
end