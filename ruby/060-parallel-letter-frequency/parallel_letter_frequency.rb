class ParallelLetterFrequency
  def self.count(texts)
    letter_freq = Hash.new(0)
    return letter_freq if texts.empty?

    threads = texts.map do |text|
      Thread.new do
        format_text = text.downcase.gsub(/[^\p{L}]/, '')
        format_text.each_char { |char| letter_freq[char] += 1 }
      end
    end
    threads.each(&:join)
    letter_freq
  end
end