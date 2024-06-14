module Port
  IDENTIFIER = :PALE

  def self.get_identifier(city)
    # city.upcase.chars.first(4).join('').to_sym
    city[0..3].upcase.to_sym
  end

  def self.get_terminal(ship_identifier)
    # type = ship_identifier.to_s.chars.first(3).join('')
    material = ship_identifier.to_s[0..2]
    if %w[OIL GAS].include?(material)
      :A
    else
      :B
    end
  end
end
