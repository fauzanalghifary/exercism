class Crypto
  # @param [String] plaintext
  def initialize(plaintext)
    @plaintext = plaintext
  end

  def ciphertext
    return "" if @plaintext.empty?

    text = @plaintext.downcase.gsub(/[^a-z0-9]/, '')
    return "" if text.empty?

    rows, cols = calculate_rectangle_dimensions(text.size)

    return "" if cols == 0

    result_chunks = []
    cols.times do |i|
      chunk = ""
      current_index = i
      rows.times do
        chunk << text[current_index] if current_index < text.length
        current_index += cols
      end
      result_chunks << chunk.ljust(rows)
    end

    result_chunks.join(" ")
  end

  def calculate_rectangle_dimensions(length)
    return [0, 0] if length == 0

    cols = Math.sqrt(length).ceil

    if cols > 0 && (cols - 1) * cols >= length
      rows = cols - 1
    else
      rows = cols
    end

    [rows, cols]
  end
end