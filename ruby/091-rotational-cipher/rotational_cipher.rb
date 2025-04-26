class RotationalCipher
  def self.rotate(text, shift)
    text.chars.map do |char|
      case char
      when 'a'..'z'
        ((((char.ord - 'a'.ord) + shift) % 26) + 'a'.ord).chr
      when 'A'..'Z'
        ((((char.ord - 'A'.ord) + shift) % 26) + 'A'.ord).chr
      else
        char
      end
    end.join
  end
end