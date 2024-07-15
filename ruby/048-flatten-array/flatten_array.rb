class FlattenArray
  def self.flatten(nested_array)
    flatten_array = []

    nested_array.each do |array|
      if !array.is_a?(Array)
        flatten_array.push(array) unless array.nil?
      else
        flatten_array.push(*flatten(array))
      end
    end

    flatten_array
  end


  # def self.flatten(array)
  #   array.flatten.compact
  # end
end