class Integer
  def to_roman
    num_to_roman = {
      "1000": "M",
      "900": "CM",
      "500": "D",
      "400": "CD",
      "100": "C",
      "90": "XC",
      "50": "L",
      "40": "XL",
      "10": "X",
      "9": "IX",
      "5": "V",
      "4": "IV",
      "1": "I",
    }

    result = ""
    remaining = self
    num_to_roman.each do |numeral|
      num = numeral[0].to_s.to_i
      roman = numeral[1]
      while remaining / num >= 1
        result << roman
        remaining -= num
      end

      break if remaining == 0
    end

    result
  end
end