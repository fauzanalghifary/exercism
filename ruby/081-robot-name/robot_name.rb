# frozen_string_literal: true

require 'set'


class Robot
  attr_reader :name

  def initialize
    reset
  end

  def reset
    @name = @@robot_names.pop
  end

  def self.forget
    @@robot_names = ('AA000'..'ZZ999').to_a.shuffle
  end
end

# class Robot
#   attr_reader :name
#   @@used_names = Set.new
#
#   def initialize
#     reset
#   end
#
#   def reset
#     @name = generate_unique_name
#   end
#
#   def self.forget
#     @@used_names.clear
#   end
#
#   private
#
#   def generate_unique_name
#     loop do
#       new_name = generate_random_name
#       return new_name if @@used_names.add?(new_name)
#     end
#   end
#
#   def generate_random_name
#     letters = Array.new(2) { ('A'..'Z').to_a.sample }.join
#     digits = format('%03d', rand(1000)) # More direct way to format 3 digits
#     "#{letters}#{digits}"
#   end
# end
#
#
