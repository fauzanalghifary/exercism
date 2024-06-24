class Hamming
  def self.compute(dna1, dna2)

    # raise ArgumentError if dna1.chars.size != dna2.chars.size
    raise ArgumentError if dna1.length != dna2.length

    # sum = 0
    #
    # dna1.chars.each_with_index do |char, index|
    #   sum += 1 if dna2[index] != char
    # end
    #
    # sum
    dna1.chars.zip(dna2.chars).count { |a, b| a != b }
  end
end