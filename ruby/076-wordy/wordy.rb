class WordProblem
  OPERATIONS = {
    'plus' => :+,
    'minus' => :-,
    'multiplied' => :*,
    'divided' => :/
  }

  def initialize(sentence)
    @numbers = []
    @operations = []

    sentence.split.each do |word|
      if (num = word[/[+-]?\d+/])
        @numbers << num.to_i
      elsif (op = OPERATIONS[word])
        @operations << op
      end
    end
  end

  def answer
    return @result if @result
    raise ArgumentError if @numbers.size != @operations.size + 1 || @operations.empty?

    result = @numbers.shift
    @operations.each do |op|
      result = result.send(op, @numbers.shift)
    end
    @result = result
  end
end
