class ResistorColorTrio
  attr_reader :colors

  COLOR_BANDS = {
    black: 0,
    brown: 1,
    red: 2,
    orange: 3,
    yellow: 4,
    green: 5,
    blue: 6,
    violet: 7,
    grey: 8,
    white: 9
  }.freeze

  def initialize(colors)
    @colors = colors
  end

  def label
    sum = 0
    color1, color2, color3 = colors
    multiplier = COLOR_BANDS[color3.to_sym]
    sum += COLOR_BANDS[color1.to_sym] * 10**(multiplier + 1) + COLOR_BANDS[color2.to_sym] * 10**multiplier
    return "Resistor value: #{sum} ohms" if sum < 1000

    "Resistor value: #{sum/1000} kiloohms"
  end
end