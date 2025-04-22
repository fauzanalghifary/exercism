class Change
  attr_reader :coins, :amount, :change_options

  def initialize(coins, amount)
    @coins = coins
    @amount = amount
    @change_options = []
  end

  def generate
    check_for_errors
    (0...coins.count).each { |start| find_combos(coins.reverse[start..-1]) }
    find_alternates if change_options.empty?
    raise ImpossibleCombinationError if change_options.empty?
    change_options.min_by(&:count)
  end

  def self.generate(coins, amount)
    Change.new(coins, amount).generate
  end

  private

  def check_for_errors
    raise NegativeTargetError if amount.negative?
    raise ImpossibleCombinationError if coins.min > amount && amount != 0
  end

  def find_combos(coin_set, total = amount, change_combo = [])
    coin_set.each do |coin|
      until coin > total
        total -= coin
        change_combo.unshift(coin)
      end
    end
    validate(change_combo, total)
  end

  def find_alternates(change_combo = [], total = amount)
    coins.each_with_index do |coin, index|
      until coins[index + 1..-1].any? { (total % _1).zero? } || total <= 0
        total -= coin
        change_combo.push(coin)
      end
    end
    validate(change_combo, total)
  end

  def validate(coins, total)
    change_options.push(coins) if total.zero?
  end

  class ImpossibleCombinationError < StandardError
  end

  class NegativeTargetError < StandardError
  end
end