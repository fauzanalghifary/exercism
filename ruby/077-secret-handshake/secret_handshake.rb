class SecretHandshake
  # @param[Numeric]num
  def initialize(num)
    @num = num
  end

  def commands
    return [] if @num.to_s.to_i != @num

    actions = []
    binary = []
    n = 16
    while n > 0
      div = @num / n
      binary << div
      @num -= n if div == 1
      n = n / 2
    end

    binary.reverse.each_with_index do |bin, index|
      next if bin != 1

      case index
      when 0 then actions << "wink"
      when 1 then actions << "double blink"
      when 2 then actions << "close your eyes"
      when 3 then actions << "jump"
      when 4 then actions.reverse!
      end
    end

    actions
  end
end