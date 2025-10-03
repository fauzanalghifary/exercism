class Yacht
  def initialize(dice, category)
    @dice = dice
    @category = category
  end

  def score
    case @category
    when 'yacht'
      yacht
    when 'ones'
      number_category(1)
    when 'twos'
      number_category(2)
    when 'threes'
      number_category(3)
    when 'fours'
      number_category(4)
    when 'fives'
      number_category(5)
    when 'sixes'
      number_category(6)
    when 'full house'
      full_house
    when 'four of a kind'
      four_of_a_kind
    when 'little straight'
      little_straight
    when 'big straight'
      big_straight
    when 'choice'
      choice
    else
      0
    end
  end

  private

  def yacht
    @dice.uniq.length == 1 ? 50 : 0
  end

  def number_category(number)
    @dice.select { |die| die == number }.sum
  end

  def full_house
    counts = @dice.tally.values.sort
    counts == [2, 3] ? @dice.sum : 0
  end

  def four_of_a_kind
    counts = @dice.tally
    die_value = counts.find { |_, count| count >= 4 }&.first
    die_value ? die_value * 4 : 0
  end

  def little_straight
    @dice.sort == [1, 2, 3, 4, 5] ? 30 : 0
  end

  def big_straight
    @dice.sort == [2, 3, 4, 5, 6] ? 30 : 0
  end

  def choice
    @dice.sum
  end
end
