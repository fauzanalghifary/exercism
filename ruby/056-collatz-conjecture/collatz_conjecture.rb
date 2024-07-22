class CollatzConjecture
  # @param [Integer] num
  def self.steps(num, count = 0)
    raise ArgumentError if num < 1
    return count if num == 1

    n = if num.even?
          num / 2
        else
          num * 3 + 1
        end
    steps(n, count + 1)
  end


  # # @param [Integer] num
  # def self.steps(num)
  #   raise ArgumentError if num < 1
  #
  #   n = num
  #   count = 0
  #   while n != 1
  #     n = if n.even?
  #           n / 2
  #         else
  #           n * 3 + 1
  #         end
  #     count += 1
  #   end
  #
  #   count
  # end
end