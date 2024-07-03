class Clock
  # attr_reader :time
  #
  # def initialize(time)
  #   @time = time
  # end
  #
  # def to_s
  #   total_minutes = time[:minute] || 0
  #   additional_hours = total_minutes / 60
  #   total_minutes %= 60
  #
  #   total_hours = (time[:hour] || 0) + additional_hours
  #   total_hours %= 24
  #
  #   minute = total_minutes.to_s.rjust(2, '0')
  #   hour = total_hours.to_s.rjust(2, '0')
  #
  #   "#{hour}:#{minute}"
  # end
  #
  # def +(other)
  #   total_minutes = (time[:minute] || 0) + (other.time[:minute] || 0)
  #   additional_hours = total_minutes / 60
  #   total_minutes %= 60
  #
  #   total_hours = (time[:hour] || 0) + (other.time[:hour] || 0) + additional_hours
  #   total_hours %= 24
  #
  #   Clock.new(hour: total_hours, minute: total_minutes)
  # end
  #
  # def -(other)
  #   total_minutes_current = (time[:minute] || 0) + ((time[:hour] || 0) * 60)
  #   total_minutes_substract = (other.time[:minute] || 0) + ((other.time[:hour] || 0) * 60)
  #   remaining_minutes = total_minutes_current - total_minutes_substract
  #
  #   Clock.new(minute: remaining_minutes)
  # end
  #
  # def ==(other)
  #   other.to_s == to_s
  # end

  attr_reader :hour, :minute

  def initialize(hour: 0, minute: 0)
    @hour = (hour + minute / 60) % 24
    @minute = minute % 60
  end

  def +(other)
    hour = self.hour + other.hour
    minute = self.minute + other.minute
    Clock.new(hour: hour, minute: minute)
  end

  def -(other)
    hour = self.hour - other.hour
    minute = self.minute - other.minute
    Clock.new(hour: hour, minute: minute)
  end

  def ==(other)
    hour == other.hour && minute == other.minute
  end

  def to_s
    format('%02d:%02d', hour, minute)
  end
end