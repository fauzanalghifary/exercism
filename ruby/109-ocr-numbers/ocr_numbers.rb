module OcrNumbers
  DIGITS = {
    [" _ ", "| |", "|_|", "   "] => "0",
    ["   ", "  |", "  |", "   "] => "1",
    [" _ ", " _|", "|_ ", "   "] => "2",
    [" _ ", " _|", " _|", "   "] => "3",
    ["   ", "|_|", "  |", "   "] => "4",
    [" _ ", "|_ ", " _|", "   "] => "5",
    [" _ ", "|_ ", "|_|", "   "] => "6",
    [" _ ", "  |", "  |", "   "] => "7",
    [" _ ", "|_|", "|_|", "   "] => "8",
    [" _ ", "|_|", " _|", "   "] => "9"
  }

  def self.convert(input)
    lines = input.split("\n")

    raise ArgumentError, "Number of input lines must be a multiple of 4" unless lines.size % 4 == 0
    raise ArgumentError, "Number of input columns must be a multiple of 3" unless lines.all? { |line| line.size % 3 == 0 }

    results = []
    lines.each_slice(4) do |group|
      results << convert_row(group)
    end

    results.join(",")
  end

  def self.convert_row(lines)
    return "" if lines.empty?

    num_digits = lines[0].size / 3
    digits = []

    num_digits.times do |i|
      start_col = i * 3
      digit_pattern = lines.map { |line| line[start_col, 3] }
      digits << (DIGITS[digit_pattern] || "?")
    end

    digits.join
  end
end
