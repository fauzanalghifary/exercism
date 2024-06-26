class BirdCount
  BUSY_LIMIT = 5

  def self.last_week
    [0, 2, 5, 3, 7, 8, 4]
  end

  def initialize(birds_per_day)
    @birds_per_day = birds_per_day
  end

  def yesterday
    @birds_per_day[-2]
  end

  def total
    @birds_per_day.sum
  end

  def busy_days
    @birds_per_day.count { |birds| birds >= BUSY_LIMIT }
  end

  def day_without_birds?
    # @birds_per_day.any?(&:zero?)
    @birds_per_day.any? {|n| n.zero?}
  end
end
