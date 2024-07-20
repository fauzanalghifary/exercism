class ETL
  # @param [Hash] old_point
  def self.transform(old_system)
    new_system = {}
    old_system.each do |point, chars|
      chars.each { |char| new_system[char.downcase] = point }
    end
    new_system.sort_by { |key, value| key }.to_h
  end
end