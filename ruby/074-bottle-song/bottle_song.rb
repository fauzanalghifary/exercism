class BottleSong
  NUM_MAP = {
    10 => 'Ten',
    9 => 'Nine',
    8 => 'Eight',
    7 => 'Seven',
    6 => 'Six',
    5 => 'Five',
    4 => 'Four',
    3 => 'Three',
    2 => 'Two',
    1 => 'One',
    0 => 'No'
  }

  # @param [Numeric] bottle_num
  # @param [Numeric] fall_bottle_num
  def self.recite(bottle_count, fall_bottle_count)
    result = ""
    remaining_bottle = bottle_count

    1.step(to: fall_bottle_count) do |i|
      result += <<~TEXT
        #{NUM_MAP[remaining_bottle]} green #{self.bottles(remaining_bottle)} hanging on the wall,
        #{NUM_MAP[remaining_bottle]} green #{self.bottles(remaining_bottle)} hanging on the wall,
        And if one green bottle should accidentally fall,
        There'll be #{NUM_MAP[remaining_bottle - 1].downcase} green #{self.bottles(remaining_bottle - 1)} hanging on the wall.
      TEXT
      remaining_bottle -= 1

      if i != fall_bottle_count
        result += <<~TEXT

        TEXT
      end
    end

    result
  end

  private

  def self.bottles(bottle_count)
    bottle_count == 1 ? 'bottle' : 'bottles'
  end
end