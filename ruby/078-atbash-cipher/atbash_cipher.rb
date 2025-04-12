class Atbash
  A_BYTE = 'a'.ord
  Z_BYTE = 'z'.ord

  # @param [String] text
  def self.encode(text)
    result = ''

    text.each_char do |char|
      c = char.downcase.ord

      if numeric?(c)
        result << char
      elsif alphabet?(c)
        result << encrypt_char(c)
      else
        next
      end

      result << ' ' if should_add_space?(result)
    end

    result.strip
  end

  # @param [String] text
  def self.decode(text)
    result = ''

    text.each_char do |char|
      c = char.ord
      next unless alphabet?(c) || numeric?(c)
      result << (numeric?(c) ? char : encrypt_char(c))
    end

    result
  end

  private

  def self.should_add_space?(text)
    text.delete(' ').size % 5 == 0
  end

  def self.encrypt_char(c)
    (Z_BYTE - c + A_BYTE).chr
  end

  def self.alphabet?(c)
    A_BYTE <= c && c <= Z_BYTE # 'a'...'z'
  end

  def self.numeric?(c)
    c.between?(48, 57) # '0'...'9'
  end
end
