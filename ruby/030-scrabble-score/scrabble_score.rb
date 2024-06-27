class Scrabble
  # LETTER_VALUES = {
  #   %w[A E I U O L N R S T] => 1,
  #   %w[D G] => 2, %w[B C M P] => 3,
  #   %w[F H V W Y] => 4, %w[K] => 5,
  #   %w[J X] => 8, %w[Q Z] => 10 }.freeze
  LETTER_VALUES = {
    /[AEIOULNRST]/ => 1,
    /[DG]/ => 2,
    /[BCMP]/ => 3,
    /[FHVWY]/ => 4,
    /[K]/ => 5,
    /[JX]/ => 8,
    /[QZ]/ => 10
  }.freeze


  def initialize(string)
    @string = string.upcase
  end

  def self.score(word)
    (new word).score
  end

  def score
    sum = 0
    LETTER_VALUES.each do |letters, value|
      sum += (@string.scan(letters).count * value)
    end
    sum
  end
end