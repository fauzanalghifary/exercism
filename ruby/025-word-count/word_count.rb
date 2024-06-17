class Phrase
  attr_reader :sentence

  def initialize(sentence)
    @sentence = sentence
  end

  def word_count
    result = {}
    sentence.split(/[^'\w]/).each do |word|
      downcase_word = word.downcase
      downcase_word = downcase_word[1..] if downcase_word[0] == "'"
      downcase_word = downcase_word[..-2] if downcase_word[-1] == "'"
      next if ['', "'"].include?(downcase_word)


      result[downcase_word] ? result[downcase_word] += 1 : result[downcase_word] = 1
    end
    result
  end
end