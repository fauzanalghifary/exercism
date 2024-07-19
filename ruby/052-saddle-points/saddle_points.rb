class Grid

  # @param [Array<Array>] matrix
  def self.saddle_points(matrix)
    result = []
    transpose_matrix = matrix.transpose
    matrix.each_with_index do |row, index_row|
      highest_tree_in_row = row.max
      row.each_with_index do |num, index_col|
        lowest_tree_in_column = transpose_matrix[index_col].min
        next unless num == highest_tree_in_row && num == lowest_tree_in_column

        result.push({
                      'row' => index_row + 1,
                      'column' => index_col + 1
                    })
      end
    end
    result
  end
end