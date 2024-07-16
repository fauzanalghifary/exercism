class Grains
  def self.square(num)
    raise ArgumentError, 'Invalid input' if num < 1 || num > 64

    2**(num - 1)
  end

  def self.total
    2**64 - 1
  end
end