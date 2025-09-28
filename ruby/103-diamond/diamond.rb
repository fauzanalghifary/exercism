class Diamond
  def self.make_diamond(letter)
    return "A\n" if letter == "A"

    diamond_size = letter.ord - "A".ord
    result = []

    # Build the top half (including the middle)
    (0..diamond_size).each do |i|
      current_letter = (65 + i).chr
      line = build_line(current_letter, diamond_size)
      result << line
    end

    # Build the bottom half (excluding the middle)
    (diamond_size - 1).downto(0) do |i|
      current_letter = (65 + i).chr
      line = build_line(current_letter, diamond_size)
      result << line
    end

    result.join("\n") + "\n"
  end

  private

  def self.build_line(letter, diamond_size)
    if letter == "A"
      return " " * diamond_size + "A" + " " * diamond_size
    end

    letter_index = letter.ord - "A".ord
    outer_spaces = " " * (diamond_size - letter_index)
    inner_spaces = " " * (2 * letter_index - 1)

    outer_spaces + letter + inner_spaces + letter + outer_spaces
  end
end