class Dominoes
  # @param [Array<Array<Integer>>]
  # @return [Boolean]
  def self.chain?(dominoes)
    return true if dominoes.empty?
    
    # Try each domino as the starting point
    dominoes.each_with_index do |first, index|
      remaining = dominoes.dup
      remaining.delete_at(index)
      
      # Try both orientations of the first domino
      [first, first.reverse].each do |oriented_first|
        # If we can build a chain starting with this domino that ends with the right number,
        # we've found a solution
        if can_chain?(remaining, oriented_first.first, oriented_first.last)
          return true
        end
      end
    end
    
    false
  end

  private

  def self.can_chain?(dominoes, start_number, current_number, used = [])
    # Base case: if we've used all dominoes, check if the chain is valid
    return start_number == current_number if dominoes.empty?
    
    # Try each remaining domino
    dominoes.each_with_index do |domino, index|
      remaining = dominoes.dup
      remaining.delete_at(index)
      
      # Try both orientations of the current domino
      if domino.first == current_number
        return true if can_chain?(remaining, start_number, domino.last, used + [domino])
      end
      
      if domino.last == current_number
        return true if can_chain?(remaining, start_number, domino.first, used + [domino.reverse])
      end
    end
    
    false
  end
end