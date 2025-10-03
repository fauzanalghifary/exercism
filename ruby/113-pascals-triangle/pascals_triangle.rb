class Triangle
  def initialize(num_rows)
    @num_rows = num_rows
  end

  def rows
    return [] if @num_rows.zero?

    triangle = []

    (0...@num_rows).each do |row_index|
      row = []
      (0..row_index).each do |col_index|
        if col_index == 0 || col_index == row_index
          row << 1
        else
          row << triangle[row_index - 1][col_index - 1] + triangle[row_index - 1][col_index]
        end
      end
      triangle << row
    end

    triangle
  end
end
