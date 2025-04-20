class PigLatin
  VOCAL = %w(a i u e o)

  # @param[String]text
  def self.translate(text)
    text.split(" ").map { |word| translate_word(word) }.join(" ")
  end

  private

  # @param[String]text
  def self.translate_word(text)
    vocal_index = 0
    text.each_char.each_with_index do |c, index|
      if text[0..1] == "xr" || text[0..1] == "yt"
        return text + "ay"
      end

      if VOCAL.include?(c)
        return text + "ay" if vocal_index == 0
        return text[vocal_index..] + text[0..vocal_index-1] + "ay"
      end

      if text[index..index+1] == "qu"
        return text[index+2..] + text[0..index+1] + "ay"
      end

      if c == "y" && index != 0
        return "y" + text[index+ 1..] + text[0..index-1] + "ay"
      end

      vocal_index += 1
    end

    text
  end
end