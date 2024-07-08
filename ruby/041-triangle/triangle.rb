class Triangle
  # attr_reader :side_a, :side_b, :side_c
  #
  # def initialize(sides)
  #   @side_a = sides[0]
  #   @side_b = sides[1]
  #   @side_c = sides[2]
  # end
  #
  # def equilateral?
  #   return false if invalid_triangle?
  #
  #   side_a == side_b && side_b == side_c
  # end
  #
  # def isosceles?
  #   return false if invalid_triangle?
  #
  #   [side_a, side_b, side_c].uniq.count <= 2
  # end
  #
  # def scalene?
  #   return false if invalid_triangle?
  #
  #   [side_a, side_b, side_c].uniq.count == 3
  # end
  #
  # def invalid_triangle?
  #   return true if side_a <= 0 || side_b <= 0 || side_c <= 0
  #
  #   true unless side_a + side_b >= side_c && side_a + side_c >= side_b && side_b + side_c >= side_a
  # end

  attr_reader :sides

  def initialize(sides)
    @sides = sides
  end

  def equilateral?
    valid_triangle? && sides.uniq.count == 1
  end

  def isosceles?
    valid_triangle? && sides.uniq.count <= 2
  end

  def scalene?
    valid_triangle? && sides.uniq.count == 3
  end

  private

  def valid_triangle?
    sides.min.positive? && sides.combination(2).all? { |a, b| a + b > sides.sum - a - b }
  end
end