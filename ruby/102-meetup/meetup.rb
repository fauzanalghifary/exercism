require 'date'

class Meetup
  # @param[Numeric]month
  # @param[Numeric]year
  def initialize(month, year)
    @month = month
    @year = year
  end

  def day(weekday, schedule)
    days = find_days_of_week(weekday)
    
    case schedule
    when :teenth
      days.find { |day| (13..19).include?(day.mday) }
    when :first
      days.first
    when :second
      days[1]
    when :third
      days[2]
    when :fourth
      days[3]
    when :last
      days.last
    end
  end

  private

  def find_days_of_week(weekday)
    day_number = day_to_number(weekday)
    
    # Create a range for all days in the month
    first_day = Date.new(@year, @month, 1)
    last_day = Date.new(@year, @month, -1)
    
    # Find all days that match the weekday
    (first_day..last_day).select { |date| date.wday == day_number }
  end

  def day_to_number(day)
    {
      monday: 1,
      tuesday: 2,
      wednesday: 3,
      thursday: 4,
      friday: 5,
      saturday: 6,
      sunday: 0
    }[day]
  end
end