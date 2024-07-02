class Luhn
  def self.valid?(input)
    # trimmed_input = input.gsub(/\s+/, '').strip
    # return false if trimmed_input.size <= 1
    #
    # doubled = trimmed_input.reverse.chars
    # doubled.each_with_index do |_, index|
    #   return false unless doubled[index].match(/\d/)
    #
    #   if index.odd?
    #     new_num = doubled[index].to_i * 2
    #     new_num -= 9 if new_num > 9
    #     doubled[index] = new_num
    #   else
    #     doubled[index] = doubled[index].to_i
    #   end
    # end
    # sum = doubled.inject(0) { |accumulator, number| accumulator + number }
    # (sum % 10).zero?
    (input
      .gsub(/\s/, '')
      .tap { |s| return false unless s[/\A\d\d+\z/] }
      .chars
      .reverse
      .map.with_index { |d, i| i.odd? ? d.to_i * 2 : d.to_i }
      .map { |d| d > 9 ? d - 9 : d }
      .sum % 10).zero?
  end
end