class Year
  def self.leap?(year)
    # if (year % 4).zero?
    #   return false if (year % 100).zero? && !(year % 400).zero?
    #
    #   return true
    # end
    # false
    year % 4 == 0 && (year % 100 != 0 || year % 400 == 0)
  end
end