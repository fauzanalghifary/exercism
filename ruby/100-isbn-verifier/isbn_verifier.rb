class IsbnVerifier
  def self.valid?(input)
    isbn = input.delete('-').upcase
    return false unless isbn.length == 10

    checksum = 0
    isbn.each_char.with_index do |char, index|
      if index < 9
        return false unless char.match?(/\d/)
        checksum += char.to_i * (10 - index)
      else
        if char == 'X'
          checksum += 10
        elsif char.match?(/\d/)
          checksum += char.to_i
        else
          return false
        end
      end
    end

    checksum % 11 == 0
  end
end