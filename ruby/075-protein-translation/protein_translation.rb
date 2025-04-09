class InvalidCodonError < RuntimeError
end

class Translation
  CODON_MAP = {
    ["AUG"] => "Methionine",
    ["UUU", "UUC"] => "Phenylalanine",
    ["UUA", "UUG"] => "Leucine",
    ["UCU", "UCC", "UCA", "UCG"] => "Serine",
    ["UAU", "UAC"] => "Tyrosine",
    ["UGU", "UGC"] => "Cysteine",
    ["UGG"] => "Tryptophan",
    ["UAA", "UAG", "UGA"] => "STOP"
  }

  # @param [String] strand
  def self.of_rna(strand)
    return [] if strand.empty?

    result = []
    i = 0
    is_stop_program = false

    while true
      current_codon = strand[i..i+2]
      is_found_acid = false
      current_result_length = result.length
      CODON_MAP.each do |codon, acid|
        break if is_found_acid || is_stop_program
        codon.each do |c|
          if c == current_codon
            if acid == "STOP"
              is_stop_program = true
              break
            end
            result << acid
            is_found_acid = true
            break
          end
        end
      end

      raise InvalidCodonError if current_result_length == result.length && !is_stop_program

      i += 3
      break if is_stop_program || i == strand.length
    end

    result
  end
end