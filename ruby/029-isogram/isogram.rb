class Isogram
  def self.isogram?(input)
    # dict = {}
    #
    # input.each_char do |char|
    #   next unless char.match(/[a-zA-Z1-9]/)
    #
    #   unless dict[char.downcase]
    #     dict[char.downcase] = true
    #     next
    #   end
    #   return false
    #
    # end
    #
    # true
    letters = input.downcase.scan(/\w/)
    letters == letters.uniq
  end
end