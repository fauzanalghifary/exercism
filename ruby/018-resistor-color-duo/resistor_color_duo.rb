class ResistorColorDuo
  COLOR_CODES = {
    'Black' => 0,
    'Brown' => 1,
    'Red' => 2,
    'Orange' => 3,
    'Yellow' => 4,
    'Green' => 5,
    'Blue' => 6,
    'Violet' => 7,
    'Grey' => 8,
    'White' => 9
  }.freeze

  def self.value(array)
    code1 = COLOR_CODES[array[0].capitalize]
    code2 = COLOR_CODES[array[1].capitalize]
    code1 * 10 + code2
  end
end

# class ResistorColorDuo
#   COLORS = %w(black brown red orange yellow
#             green blue violet grey white)
#   def self.value(array)
#     array.take(2).map{|color| COLORS.index(color) }.join.to_i
#   end
# end
