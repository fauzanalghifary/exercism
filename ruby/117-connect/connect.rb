class Board
  def initialize(board)
    @board = board
  end

  def winner
    return 'X' if x_wins?
    return 'O' if o_wins?
    ''
  end

  private

  def x_wins?
    # X wins by connecting left to right
    leftmost_x_positions.any? { |pos| can_reach_right?(pos, 'X') }
  end

  def o_wins?
    # O wins by connecting top to bottom
    topmost_o_positions.any? { |pos| can_reach_bottom?(pos, 'O') }
  end

  def leftmost_x_positions
    positions = []
    @board.each_with_index do |row, r|
      # The left edge is column 0 (after gsub removes leading spaces)
      positions << [r, 0] if cell_at(r, 0) == 'X'
    end
    positions
  end

  def topmost_o_positions
    return [] if @board.empty?
    positions = []
    @board[0].chars.each_with_index do |cell, c|
      positions << [0, c] if cell == 'O'
    end
    positions
  end

  def can_reach_right?(start_pos, player)
    visited = {}
    queue = [start_pos]
    visited[start_pos] = true

    until queue.empty?
      row, col = queue.shift

      # Check if we've reached the rightmost column
      if col == rightmost_col(row)
        return true
      end

      neighbors(row, col).each do |nr, nc|
        next if visited[[nr, nc]]
        next unless cell_at(nr, nc) == player

        visited[[nr, nc]] = true
        queue << [nr, nc]
      end
    end

    false
  end

  def can_reach_bottom?(start_pos, player)
    visited = {}
    queue = [start_pos]
    visited[start_pos] = true

    until queue.empty?
      row, col = queue.shift

      return true if row == @board.length - 1

      neighbors(row, col).each do |nr, nc|
        next if visited[[nr, nc]]
        next unless cell_at(nr, nc) == player

        visited[[nr, nc]] = true
        queue << [nr, nc]
      end
    end

    false
  end

  def neighbors(row, col)
    # Hex neighbors, accounting for spaces in representation
    # a space separates Each cell, so cells are at even columns (0,2,4,6, ...)
    # After gsub, rows are left-aligned, so we need offset-aware neighbors
    [
      [row - 1, col],      # upper-left
      [row - 1, col + 2],  # upper-right
      [row, col - 2],      # left
      [row, col + 2],      # right
      [row + 1, col - 2],  # lower-left
      [row + 1, col]       # lower-right
    ].select { |r, c| valid_position?(r, c) && cell_at(r, c) != ' ' }
  end

  def valid_position?(row, col)
    row >= 0 && row < @board.length && col >= 0 && col < row_length(row)
  end

  def cell_at(row, col)
    return nil unless valid_position?(row, col)
    @board[row][col]
  end

  def row_length(row)
    @board[row].length
  end

  def rightmost_col(row)
    # Find the rightmost non-space character
    @board[row].chars.each_with_index.reverse_each do |char, idx|
      return idx if char != ' '
    end
    0
  end
end
