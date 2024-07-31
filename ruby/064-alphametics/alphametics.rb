# class Alphametics
#   # @param [String] puzzle
#   def self.solve(puzzle)
#     question, answer = puzzle.split('==').map(&:strip)
#     questions = question.split('+').map(&:strip)
#     letters = (questions.join + answer).chars.uniq
#     digits = (0..9).to_a
#
#     permutations = digits.permutation(letters.size).to_a
#
#     permutations.each do |perm|
#       mapping = letters.zip(perm).to_h
#       return mapping if valid?(questions, answer, mapping)
#     end
#
#     {}
#   end
#
#
#   def self.valid?(questions, answer, mapping)
#     question_num = questions.map { |q| substitute(q, mapping).to_i }
#     question_sum = question_num.sum
#
#     return false if questions.join('').size != question_num.join('').size
#
#     answer_sum = substitute(answer, mapping).to_i
#
#     return false if answer.size != answer_sum.to_s.size
#
#     question_sum == answer_sum
#   end
#
#   def self.substitute(word, mapping)
#     word.chars.map { |char| mapping[char] }.join
#   end
# end

### Community Solution
class Alphametics
  def self.solve(equation)
    letters = equation.scan(/\w/).uniq
    not_zero_expr = equation.scan(/\b(.)\w/).flatten.uniq.map { |char| "#{char.downcase} != 0 && " }.join

    math_expr = equation.downcase.gsub(/\w+/) do |word|
      "(#{word.chars.reverse.map.with_index { |char, i| "#{char}*#{10**i}" }.join '+'})"
    end

    checker = eval "->#{letters.join(',').downcase}{ #{not_zero_expr}#{math_expr} }"
    (0..9).to_a.permutation(letters.size) do |numbers|
      return letters.zip(numbers).to_h if checker[*numbers]
    end

    {}
  end
end