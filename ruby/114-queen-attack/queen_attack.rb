class Queens
  def initialize(white: [0, 3], black: [7, 3])
    @white = white
    @black = black

    validate_position(white, "white")
    validate_position(black, "black")
  end

  def attack?
    same_row? || same_column? || same_diagonal?
  end

  private

  def validate_position(position, color)
    row, col = position

    if row < 0 || row > 7
      raise ArgumentError, "#{color} queen must have row on board"
    end

    if col < 0 || col > 7
      raise ArgumentError, "#{color} queen must have column on board"
    end
  end

  def same_row?
    @white[0] == @black[0]
  end

  def same_column?
    @white[1] == @black[1]
  end

  def same_diagonal?
    row_diff = (@white[0] - @black[0]).abs
    col_diff = (@white[1] - @black[1]).abs
    row_diff == col_diff
  end
end
