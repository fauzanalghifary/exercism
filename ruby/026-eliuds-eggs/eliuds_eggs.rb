class EliudsEggs
  def self.egg_count(display_num)
    array = []
    num = display_num
    while num.positive?
      array.unshift(num % 2)
      num /= 2
    end

    array.count { |number| number == 1 }

    # display_num.to_s(2).count('1')
  end
end