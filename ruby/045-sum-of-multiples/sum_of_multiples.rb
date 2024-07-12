class SumOfMultiples
  attr_reader :nums

  def initialize(*nums)
    @nums = nums
  end


  def to(level)
    array1 = []

    (1...level).each do |target_num|
      nums.each do |num|
        if (target_num % num).zero?
          array1.push(target_num)
          break
        end
      end
    end

    array1.sum

    # (1...level).select do|target_num|
    #   @nums.any? { |num| target_num % num == 0 }
    # end.sum
  end
end