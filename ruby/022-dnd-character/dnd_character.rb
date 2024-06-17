class DndCharacter
  attr_reader :strength, :dexterity, :constitution, :intelligence, :wisdom, :charisma

  def initialize
    @strength = points
    @dexterity = points
    @constitution = points
    @intelligence = points
    @wisdom = points
    @charisma = points
  end

  def hitpoints
    10 + DndCharacter.modifier(constitution)
    # 10 + self.class.modifier(constitution)
  end

  def self.modifier(constitution)
    ((constitution - 10) / 2).floor
  end

  def points
    arr = []
    4.times { arr << rand(1..6) }
    arr.sort.reverse.take(3).sum
  end
end
