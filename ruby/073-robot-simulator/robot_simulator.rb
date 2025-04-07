class Robot
  DIRECTIONS = %i[north east south west]

  def orient(dir)
    raise ArgumentError unless DIRECTIONS.include?(dir)
    @direction = dir
  end

  def bearing
    @direction
  end

  def at(x, y)
    @x = x
    @y = y
  end

  def coordinates
    [@x, @y]
  end

  def advance
    case @direction
    when :north then @y += 1
    when :east  then @x += 1
    when :south then @y -= 1
    when :west  then @x -= 1
    end
  end

  def turn_right
    rotate(1)
  end

  def turn_left
    rotate(-1)
  end

  private

  def rotate(offset)
    idx = DIRECTIONS.index(@direction)
    @direction = DIRECTIONS[(idx + offset) % 4]
  end
end


class Simulator
  INSTRUCTION_MAP = {
    'L' => :turn_left,
    'R' => :turn_right,
    'A' => :advance
  }

  def instructions(command)
    command.chars.map { |char| INSTRUCTION_MAP[char] }
  end

  # @param [Robot] robot
  def place(robot, x:, y:, direction:)
    robot.orient(direction)
    robot.at(x, y)
  end

  # @param [Robot] robot
  def evaluate(robot, commands)
    instructions(commands).each { |cmd| robot.send(cmd) }
  end
end