class Say
  attr_reader :number

  LESS_THAN_20 = %w[zero one two three four five six seven eight nine ten eleven twelve thirteen fourteen fifteen sixteen seventeen eighteen nineteen].freeze
  TENS = %w[_ _ twenty thirty forty fifty sixty seventy eighty ninety].freeze
  SCALES = %w[thousand million billion].freeze

  def initialize(number)
    raise ArgumentError, "Number must be between 0 and 999,999,999,999" unless (0..999_999_999_999).cover?(number)
    @number = number
  end

  def in_english
    return LESS_THAN_20[0] if number.zero?

    parts = []
    num = number
    scale_index = -1 # Start before 'thousand'

    while num > 0
      num, remainder = num.divmod(1000)
      if remainder > 0
        scale_word = scale_index >= 0 ? SCALES[scale_index] : nil
        parts << "#{process_hundreds(remainder)}#{" " + scale_word if scale_word}"
      end
      scale_index += 1
    end

    parts.reverse.join(' ')
  end

  private

  def process_hundreds(n)
    parts = []
    hundreds, remainder = n.divmod(100)

    if hundreds > 0
      parts << "#{LESS_THAN_20[hundreds]} hundred"
    end

    if remainder > 0
      parts << process_tens(remainder)
    end

    parts.join(' ')
  end

  def process_tens(n)
    if n < 20
      LESS_THAN_20[n]
    else
      tens, ones = n.divmod(10)
      result = TENS[tens]
      result += "-#{LESS_THAN_20[ones]}" if ones > 0
      result
    end
  end
end