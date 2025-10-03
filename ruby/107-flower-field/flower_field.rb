class FlowerField
  def self.annotate(input)
    return input if input.empty? || input.all?(&:empty?)

    rows = input.size
    cols = input[0].size
    result = input.map(&:dup)

    rows.times do |row|
      cols.times do |col|
        next if input[row][col] == '*'

        count = count_adjacent_flowers(input, row, col, rows, cols)
        result[row][col] = count > 0 ? count.to_s : ' '
      end
    end

    result
  end

  private

  def self.count_adjacent_flowers(input, row, col, rows, cols)
    count = 0

    (-1..1).each do |dr|
      (-1..1).each do |dc|
        next if dr == 0 && dc == 0

        nr, nc = row + dr, col + dc
        next if nr < 0 || nr >= rows || nc < 0 || nc >= cols

        count += 1 if input[nr][nc] == '*'
      end
    end

    count
  end
end
