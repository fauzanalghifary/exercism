class Game
  class BowlingError < StandardError; end

  def initialize
    @rolls = []
  end

  def roll(pins)
    raise BowlingError, "Pins must be between 0 and 10" if pins < 0 || pins > 10
    raise BowlingError, "Cannot roll after game is over" if game_complete?

    if in_last_frame?
      validate_last_frame_roll(pins)
    else
      validate_normal_frame_roll(pins)
    end

    @rolls << pins
  end

  def score
    raise BowlingError, "Game is not complete" unless game_complete?

    total_score = 0
    roll_index = 0

    10.times do |frame|
      if strike?(roll_index)
        total_score += 10 + @rolls[roll_index + 1] + @rolls[roll_index + 2]
        roll_index += 1
      elsif spare?(roll_index)
        total_score += 10 + @rolls[roll_index + 2]
        roll_index += 2
      else
        total_score += @rolls[roll_index] + @rolls[roll_index + 1]
        roll_index += 2
      end
    end

    total_score
  end

  private

  def strike?(roll_index)
    @rolls[roll_index] == 10
  end

  def spare?(roll_index)
    @rolls[roll_index] + @rolls[roll_index + 1] == 10
  end

  def current_frame
    frame = 0
    roll_index = 0

    while roll_index < @rolls.size && frame < 10
      if frame == 9
        # We're in the last frame
        return 9
      end

      if @rolls[roll_index] == 10
        roll_index += 1
        frame += 1
      elsif roll_index + 1 < @rolls.size
        roll_index += 2
        frame += 1
      else
        # We're in the middle of a frame
        return frame
      end
    end

    frame
  end

  def in_last_frame?
    current_frame == 9
  end

  def rolls_in_current_frame
    frame = 0
    roll_index = 0
    target_frame = current_frame

    while roll_index < @rolls.size && frame < target_frame
      if @rolls[roll_index] == 10
        roll_index += 1
      else
        roll_index += 2
      end
      frame += 1
    end

    @rolls[roll_index..-1] || []
  end

  def validate_normal_frame_roll(pins)
    frame_rolls = rolls_in_current_frame

    # If this is the second roll of a frame and first wasn't a strike
    if frame_rolls.size == 1 && frame_rolls[0] != 10
      raise BowlingError, "Pin count exceeds pins on the lane" if frame_rolls[0] + pins > 10
    end
  end

  def validate_last_frame_roll(pins)
    frame_rolls = rolls_in_current_frame

    case frame_rolls.size
    when 0
      # First roll - any valid pin count is OK
    when 1
      # Second roll
      if frame_rolls[0] != 10
        # First roll wasn't a strike, so total can't exceed 10
        raise BowlingError, "Pin count exceeds pins on the lane" if frame_rolls[0] + pins > 10
      end
      # If first was a strike, any count 0-10 is valid
    when 2
      # Third roll
      if frame_rolls[0] == 10
        # First was a strike
        if frame_rolls[1] == 10
          # Second was also a strike, any count 0-10 is valid
        else
          # Second wasn't a strike, so rolls 2+3 can't exceed 10
          raise BowlingError, "Pin count exceeds pins on the lane" if frame_rolls[1] + pins > 10
        end
      else
        # First wasn't a strike (must be spare to get here), any count is valid
      end
    end
  end

  def game_complete?
    return false if @rolls.empty?

    frame = 0
    roll_index = 0

    while frame < 10 && roll_index < @rolls.size
      if frame == 9
        # Last frame
        if @rolls[roll_index] == 10
          # Strike in last frame needs 2 more rolls
          return roll_index + 2 < @rolls.size
        elsif roll_index + 1 < @rolls.size && @rolls[roll_index] + @rolls[roll_index + 1] == 10
          # Spare in last frame needs 1 more roll
          return roll_index + 2 < @rolls.size
        else
          # No strike or spare, needs 2 rolls total
          return roll_index + 1 < @rolls.size
        end
      else
        if @rolls[roll_index] == 10
          roll_index += 1
        else
          roll_index += 2
        end
        frame += 1
      end
    end

    frame == 10
  end
end
