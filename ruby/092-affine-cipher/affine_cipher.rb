class Affine
  # @param[Numeric]a
  # @param[Numeric]b
  def initialize(a, b)
    validate_input(a)
    @a = a
    @b = b
    @mmi = mmi(a)
  end

  # @param[String]plaintext
  def encode(plaintext)
    encoded_chars = []

    plaintext.each_char do |c|
      if c.match?(/[a-z]/i)
        i = c.downcase.ord - 'a'.ord
        e = (@a * i + @b) % 26
        encoded_chars << ('a'.ord + e).chr
      elsif c.match?(/[0-9]/)
        encoded_chars << c
      end
    end

    encoded_chars.each_slice(5).map(&:join).join(' ')
  end

  # @param[String]text
  def decode(text)
    decoded_chars = []

    text.each_char do |c|
      if c.match?(/[a-z]/i) 
        char_num = c.downcase.ord - 'a'.ord
        decrypted_num = (@mmi * (char_num - @b) % 26 + 26) % 26 
        decoded_chars << ('a'.ord + decrypted_num).chr
      elsif c.match?(/[0-9]/) 
        decoded_chars << c
      end
    end

    decoded_chars.join('') 
  end

  private

  # @param[Numeric]a
  def validate_input(a)
    raise ArgumentError if gcd(26, a) != 1
  end

  def gcd(a, b)
    while b != 0
      a, b = b, a % b
    end

    a
  end

  def mmi(a)
    a = a % 26
    (1..26).find { |i| (a * i) % 26 == 1 }
  end
end