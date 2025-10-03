class Darts
  def initialize(x, y)
    @x = x
    @y = y
  end

  def score
    distance = Math.sqrt(@x**2 + @y**2)

    if distance <= 1
      10
    elsif distance <= 5
      5
    elsif distance <= 10
      1
    else
      0
    end
  end
end
