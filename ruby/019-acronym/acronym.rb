class Acronym
  def self.abbreviate(str)
    str.scan(/\b\w/).join.upcase
  end
end


# class Acronym
#   def self.abbreviate(sentence)
#     final = ''
#     words = sentence.gsub('-', ' ').upcase.split(' ')
#     words.each do |word|
#       final += word[0]
#     end
#     final
#   end
# end