class RailFenceCipher
  def self.encode(str, rails_num)
    return str if rails_num == 1

    collection = Array.new(rails_num) { "" }
    i = 0
    direction = 1

    str.each_char do |char|
      collection[i] << char
      i += direction
      direction *= -1 if i == 0 || i == rails_num - 1
    end

    collection.join
  end

  def self.decode(str, rails_num)
    return str if rails_num == 1

    rails = Array.new(rails_num) { "" }
    positions = Array.new(str.length)

    i = 0
    direction = 1
    str.length.times do |index|
      positions[index] = i
      i += direction
      direction *= -1 if i == 0 || i == rails_num - 1
    end

    pos_index = 0
    rails.each_with_index do |rail, rail_index|
      positions.each do |pos|
        if rail_index == pos
          rail << str[pos_index]
          pos_index += 1
        end
      end
    end

    result = ""
    positions.each do |rail_index|
      result << rails[rail_index].slice!(0)
    end

    result
  end
end