class Poker
  CARD_VALUE = { 'A' => 14, 'K' => 13, 'Q' => 12, 'J' => 11 }
  attr_reader :best_hand

  def initialize(hands)
    scores = hands.map { |hand| tally hand }
    max_score = scores.max
    @best_hand = hands.each_with_index.select { |hand, i| scores[i] == max_score }.map(&:first)
  end

  def tally(cards)
    @values = cards.map { |card| CARD_VALUE[card[0..-2]] || card[0..-2].to_i }.sort
    @suits = cards.map { |card| card[-1] }
    score = @values.reverse.inject(1) { |sum, value| sum*10 + value }
    score += find_sets_of(2) * 10**6
    score += find_sets_of(3) * 10**8
    score += find_straight * 10**9
    score += find_flush * 10**10
    score += find_full * 10**11
    score += find_sets_of(4) * 10**13
    score += find_straight_flush * 10**14
  end

  def find_sets_of(multiple)
    sets = @values.uniq.select { |uniq_value| @values.select { |value| value == uniq_value }.count == multiple }
    (sets[0] || 0) + (sets[1] || 0)*10
  end

  def find_straight
    values = @values[-1] == 14 && @values[0] == 2 ? [1] + @values[0..3] : @values
    (values.min..values.max).to_a == values.sort ? values.max : 0
  end

  def find_flush
    @suits.uniq.count == 1 ? @values.max : 0
  end

  def find_full
    set_of_3 = find_sets_of 3
    pair = find_sets_of 2
    set_of_3 != 0 && pair != 0 ? set_of_3 * 10 + pair : 0
  end

  def find_straight_flush
    find_flush != 0 ? find_straight : 0
  end
end

# class Hand
#   include Comparable
#
#   CARD_VALUES = {
#     '2' => 2, '3' => 3, '4' => 4, '5' => 5, '6' => 6, '7' => 7, '8' => 8,
#     '9' => 9, '10' => 10, 'J' => 11, 'Q' => 12, 'K' => 13, 'A' => 14
#   }.freeze
#
#   HAND_RANKS = {
#     high_card: 0,
#     one_pair: 1,
#     two_pair: 2,
#     three_of_a_kind: 3,
#     straight: 4,
#     flush: 5,
#     full_house: 6,
#     four_of_a_kind: 7,
#     straight_flush: 8
#   }.freeze
#
#   attr_reader :hand
#
#   def initialize(hand)
#     @hand = hand
#     @values = hand.map { |card| card[0..-2] }
#     @suits = hand.map { |card| card[-1] }
#     @value_counts = @values.group_by(&:itself).transform_values(&:count)
#   end
#
#   def <=>(other)
#     rank_comparison = rank <=> other.rank
#     return rank_comparison unless rank_comparison.zero?
#
#     # If ranks are equal, compare based on specific hand type
#     case HAND_RANKS.key(rank)
#     when :high_card, :flush
#       compare_high_cards(other)
#     when :straight, :straight_flush
#       compare_straight(other)
#     when :one_pair
#       compare_one_pair(other)
#     when :two_pair
#       compare_two_pairs(other)
#     when :three_of_a_kind
#       compare_three_of_a_kind(other)
#     when :full_house
#       compare_full_house(other)
#     when :four_of_a_kind
#       compare_four_of_a_kind(other)
#     else
#       0
#     end
#   end
#
#   def rank
#     return HAND_RANKS[:straight_flush] if straight? && flush?
#     return HAND_RANKS[:four_of_a_kind] if four_of_a_kind?
#     return HAND_RANKS[:full_house] if full_house?
#     return HAND_RANKS[:flush] if flush?
#     return HAND_RANKS[:straight] if straight?
#     return HAND_RANKS[:three_of_a_kind] if three_of_a_kind?
#     return HAND_RANKS[:two_pair] if two_pair?
#     return HAND_RANKS[:one_pair] if one_pair?
#     HAND_RANKS[:high_card]
#   end
#
#   private
#
#   def card_values
#     @card_values ||= @values.map { |v| CARD_VALUES[v] }
#   end
#
#   def sorted_values
#     @sorted_values ||= card_values.sort.reverse
#   end
#
#   def straight?
#     values = sorted_values
#     return true if values == [14, 5, 4, 3, 2] # Ace-low straight
#     values.each_cons(2).all? { |a, b| a == b + 1 }
#   end
#
#   def straight_value
#     values = sorted_values
#     return 5 if values == [14, 5, 4, 3, 2] # Ace-low straight is lowest
#     values.max
#   end
#
#   def flush?
#     @suits.uniq.size == 1
#   end
#
#   def four_of_a_kind?
#     @value_counts.values.include?(4)
#   end
#
#   def full_house?
#     @value_counts.values.sort == [2, 3]
#   end
#
#   def three_of_a_kind?
#     @value_counts.values.include?(3)
#   end
#
#   def two_pair?
#     @value_counts.values.sort == [1, 2, 2]
#   end
#
#   def one_pair?
#     @value_counts.values.sort == [1, 1, 1, 2]
#   end
#
#   def compare_high_cards(other)
#     sorted_values.zip(other.send(:sorted_values)).each do |a, b|
#       comparison = a <=> b
#       return comparison unless comparison.zero?
#     end
#     0
#   end
#
#   def compare_straight(other)
#     straight_value <=> other.send(:straight_value)
#   end
#
#   def compare_one_pair(other)
#     my_pair = @value_counts.key(2)
#     other_pair = other.instance_variable_get(:@value_counts).key(2)
#
#     pair_comparison = CARD_VALUES[my_pair] <=> CARD_VALUES[other_pair]
#     return pair_comparison unless pair_comparison.zero?
#
#     # Compare kickers
#     my_kickers = sorted_values.reject { |v| v == CARD_VALUES[my_pair] }
#     other_kickers = other.send(:sorted_values).reject { |v| v == CARD_VALUES[other_pair] }
#     my_kickers.zip(other_kickers).each do |a, b|
#       comparison = a <=> b
#       return comparison unless comparison.zero?
#     end
#     0
#   end
#
#   def compare_two_pairs(other)
#     my_pairs = @value_counts.select { |_, v| v == 2 }.keys.map { |v| CARD_VALUES[v] }.sort.reverse
#     other_pairs = other.instance_variable_get(:@value_counts).select { |_, v| v == 2 }.keys.map { |v| CARD_VALUES[v] }.sort.reverse
#
#     my_pairs.zip(other_pairs).each do |a, b|
#       comparison = a <=> b
#       return comparison unless comparison.zero?
#     end
#
#     my_kicker = card_values.find { |v| @value_counts[@values.find { |val| CARD_VALUES[val] == v }] == 1 }
#     other_kicker = other.send(:card_values).find { |v| other.instance_variable_get(:@value_counts)[other.instance_variable_get(:@values).find { |val| CARD_VALUES[val] == v }] == 1 }
#     my_kicker <=> other_kicker
#   end
#
#   def compare_three_of_a_kind(other)
#     my_three = @value_counts.key(3)
#     other_three = other.instance_variable_get(:@value_counts).key(3)
#     three_comparison = CARD_VALUES[my_three] <=> CARD_VALUES[other_three]
#     return three_comparison unless three_comparison.zero?
#
#     # Compare kickers
#     my_kickers = sorted_values.reject { |v| v == CARD_VALUES[my_three] }.sort.reverse
#     other_kickers = other.send(:sorted_values).reject { |v| v == CARD_VALUES[other_three] }.sort.reverse
#     my_kickers.zip(other_kickers).each do |a, b|
#       comparison = a <=> b
#       return comparison unless comparison.zero?
#     end
#     0
#   end
#
#   def compare_full_house(other)
#     my_three = @value_counts.key(3)
#     other_three = other.instance_variable_get(:@value_counts).key(3)
#     three_comparison = CARD_VALUES[my_three] <=> CARD_VALUES[other_three]
#     return three_comparison unless three_comparison.zero?
#
#     my_pair = @value_counts.key(2)
#     other_pair = other.instance_variable_get(:@value_counts).key(2)
#     CARD_VALUES[my_pair] <=> CARD_VALUES[other_pair]
#   end
#
#   def compare_four_of_a_kind(other)
#     my_four = @value_counts.key(4)
#     other_four = other.instance_variable_get(:@value_counts).key(4)
#     four_comparison = CARD_VALUES[my_four] <=> CARD_VALUES[other_four]
#     return four_comparison unless four_comparison.zero?
#
#     # Compare kickers
#     my_kicker = card_values.find { |v| @value_counts[@values.find { |val| CARD_VALUES[val] == v }] == 1 }
#     other_kicker = other.send(:card_values).find { |v| other.instance_variable_get(:@value_counts)[other.instance_variable_get(:@values).find { |val| CARD_VALUES[val] == v }] == 1 }
#     my_kicker <=> other_kicker
#   end
# end
#
# class Poker
#   def initialize(hands)
#     @hands = hands.map { |hand| Hand.new(hand) }
#   end
#
#   def best_hand
#     max_hand = @hands.max
#     winning_hands = @hands.select { |hand| hand == max_hand }
#     winning_hands.map(&:hand)
#   end
# end