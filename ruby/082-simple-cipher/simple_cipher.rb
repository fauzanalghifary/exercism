class Cipher
  # @param[String]text
  def initialize(text = nil)
    return @key if @key
    if text
      raise ArgumentError if text.empty?
      invalid?(text)
      @key = text
    else
      @key = ('a'..'z').to_a.sample(100).join
    end
  end

  # @param[String]plaintext
  def encode(plaintext)
    result = ''
    plaintext.each_char.each_with_index do |c, index|
      shift = @key[index].ord - 'a'.ord
      new_ord = ((c.ord - 'a'.ord + shift) % 26) + 'a'.ord
      result << new_ord.chr
    end
    result
  end

  # @param[String]text
  def decode(text)
    result = ''
    text.each_char.each_with_index do |c, index|
      shift = @key[index].ord - 'a'.ord
      new_ord = ((c.ord - 'a'.ord - shift) % 26) + 'a'.ord
      result << new_ord.chr
    end
    result
  end

  def key
    @key
  end

  private

  def invalid?(text)
    raise ArgumentError if text =~ /[^a-z]/
  end
end