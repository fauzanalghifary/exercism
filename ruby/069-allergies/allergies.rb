class Allergies
  attr_reader :allergy_score, :allergic_dict

  # @param [Numeric] allergy_score
  def initialize(allergy_score)
    @allergy_score = allergy_score
    @allergic_dict = {
      1 => 'eggs',
      2 => 'peanuts',
      4 => 'shellfish',
      8 => 'strawberries',
      16 => 'tomatoes',
      32 => 'chocolate',
      64 => 'pollen',
      128 => 'cats'
    }
  end

  def list
    array = []
    score_left = allergy_score % 256
    allergic_dict.keys.reverse.each do |num|
      if score_left >= num
        array << allergic_dict[num]
        score_left -= num
      end
    end

    array.reverse
  end

  # @param [String] food
  def allergic_to?(food)
    list.include?(food)
  end
end



## Community Solution


# class Allergies
#   ALLERGIES = %w{
#     eggs
#     peanuts
#     shellfish
#     strawberries
#     tomatoes
#     chocolate
#     pollen
#     cats
#   }.freeze
#
#   private_constant :ALLERGIES
#   def initialize(score)
#     @score = score.digits(2)
#   end
#
#   def allergic_to?(allergy)
#     @score[ALLERGIES.index(allergy)] == 1
#   end
#
#   def list
#     ALLERGIES.select(&method(:allergic_to?))
#   end
# end
