class Nucleotide
  # @param [String] dna
  # @return [Array]
  def self.from_dna(dna)
    raise ArgumentError if dna =~ /[^ATCG]/

    dna.chars
  end
end

class Array
  def histogram
    dna = {
      "A": 0,
      "T": 0,
      "C": 0,
      "G": 0
    }

    each do |item|
      dna[item.to_sym] += 1
    end
    dna.transform_keys(&:to_s)

  end
end


# class Nucleotide
#
#   def initialize(strand)
#     raise ArgumentError if strand =~ /[^ATCG]/
#     @strand = strand
#   end
#
#   def self.from_dna(strand)
#     new(strand)
#   end
#
#   def histogram
#     base_histogram = { 'A' => 0, 'T' => 0, 'C' => 0, 'G' => 0 }
#     @strand.chars.each_with_object(base_histogram) { |i, acc| acc[i] += 1 }
#   end
#
#   def count(nucleotide)
#     histogram[nucleotide]
#   end
#
# end
