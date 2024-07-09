class Anagram
  attr_reader :word

  def initialize(word)
    @word = word
  end

  # @param [Array] candidates
  def match(candidates)
    candidates.select do |target|
      next if target.size != word.size || target.downcase == word.downcase

      target.downcase.each_char.sort == word.downcase.each_char.sort
    end

  end
end
