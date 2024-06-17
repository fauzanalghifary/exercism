class Series
  def initialize(series)
    @series = series
  end

  def slices(slices_num)
    raise ArgumentError if slices_num > @series.size || slices_num <= 0

    @series                   # assuming n == 2 :
      .each_char              # %w(s t r i n g)
      .each_cons(slices_num)  # [%w(s t), %w(t r), %w(r i), %w(i n), %w(n g)]
      .map(&:join)            # %w(st tr ri in ng)

    # result = []
    #
    # index = 0
    # while index <= @series.size - slices_num
    #   result.push(@series.slice(index..index + slices_num - 1))
    #   index += 1
    # end
    #
    # result
  end
end