class Proverb
  attr_reader :input, :qualifier

  def initialize(*input, qualifier: '')
    @input = input
    @qualifier = qualifier
  end

  def to_s
    input.each_cons(2)
         .map { |lacked_item, jeopardised_item| "For want of a #{lacked_item} the #{jeopardised_item} was lost." }
         .push("And all for the want of a #{qualifier} #{input.first}.".squeeze(' '))
         .join("\n")
  end
end