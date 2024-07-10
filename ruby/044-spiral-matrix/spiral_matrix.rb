class SpiralMatrix
  attr_reader :num

  def initialize(num)
    @num = num
  end

  def matrix
    spiral = self.create_matrix(2, num)
    cursor_x = 0
    cursor_y = 0
    direction = 'right'

    (1..num**2).each do |digit|
      spiral[cursor_y][cursor_x] = digit

      case direction
      when 'right'
        if cursor_x + 1 == num || spiral[cursor_y][cursor_x + 1] != 0
          direction = 'down'
          cursor_y += 1
          next
        else
          cursor_x += 1
        end
      when 'down'
        if cursor_y + 1 == num || spiral[cursor_y + 1][cursor_x] != 0
          direction = 'left'
          cursor_x -= 1
          next
        else
          cursor_y += 1
        end
      when 'left'
        if cursor_x == 0 || spiral[cursor_y][cursor_x - 1] != 0
          p 'here move top from left', digit
          direction = 'top'
          cursor_y -= 1
          next
        else
          cursor_x -= 1
        end
      when 'top'
        if cursor_y == 0 || spiral[cursor_y - 1][cursor_x] != 0
          direction = 'right'
          cursor_x += 1
          next
        else
          cursor_y -= 1
        end
      end
    end

    spiral
  end

  def create_matrix(dimensions, size, value = 0)
    return Array.new(size) { value } if dimensions == 1

    Array.new(size) { create_matrix(dimensions - 1, size, value) }
  end
end