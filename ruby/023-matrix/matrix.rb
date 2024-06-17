class Matrix
  attr_reader :matrix

  def initialize(string)
    @matrix = string.split(/\n/).map { |row| row.split(' ').map(&:to_i) }
  end

  def row(row_num)
    matrix[row_num - 1]
  end

  def column(column_num)
    matrix.transpose[column_num - 1]
    #   @matrix.map { |row| row[position - 1] }
  end
end