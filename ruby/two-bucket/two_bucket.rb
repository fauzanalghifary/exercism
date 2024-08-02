class TwoBucket
  attr_reader :bucket_one, :bucket_two, :volume_goal, :first_fill

  # @param [Numeric] bucket_one
  # @param [Numeric] bucket_two
  # @param [Numeric] volume_goal
  # @param [String] first_fill
  def initialize(bucket_one, bucket_two, volume_goal, first_fill)
    @bucket_one = bucket_one
    @bucket_two = bucket_two
    @volume_goal = volume_goal
    @first_fill = first_fill
  end

  def moves

  end

  def goal_bucket

  end

  def other_buket

  end
end