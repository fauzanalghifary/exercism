class Pangram
  IS_ALPHABET = /[a-z]/i.freeze
  def self.pangram?(sentence)
    dict = {}
    sentence.each_char do |char|
      next unless char.match(IS_ALPHABET)

      dict[char.downcase] = char.downcase
    end
    dict.size == 26
    #  ('a'..'z').all? { |e| phrase.downcase.include?(e) }

    #  phrase.downcase.scan(/[a-z]/).uniq.size == 26
  end
end