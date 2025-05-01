class CustomSet
  def initialize(values)
    @values = values.uniq
  end

  def empty?
    values.empty?
  end

  def member?(element)
    values.include?(element)
  end

  def subset?(other)
    difference(other).empty?
  end

  def disjoint?(other)
    intersection(other).empty?
  end

  def ==(other)
    return false unless other.kind_of?(CustomSet)
    values.sort == other.values.sort
  end

  def add(element)
    values << element unless values.include?(element)
    self
  end

  def intersection(other)
    CustomSet.new(values & other.values)
  end

  def difference(other)
    CustomSet.new(values - other.values)
  end

  def union(other)
    CustomSet.new(values + other.values)
  end

  protected
  attr_reader :values
end

# class CustomSet
#   attr_reader :set
#
#   # @param[Array<Numeric>]arr
#   def initialize(arr)
#     @set = arr
#   end
#
#   def empty?
#     set.length == 0
#   end
#
#   def member?(num)
#     set.include?(num)
#   end
#
#   def subset?(other_set)
#     return false if set.length > other_set.length
#     set.each do |s|
#       return false unless other_set.member?(s)
#     end
#     true
#   end
#
#   def disjoint?(other_set)
#     set.each do |s|
#       return false if other_set.member?(s)
#     end
#     true
#   end
#
#   def intersection(other_set)
#     intersects = []
#     set.each do |s|
#       intersects << s if other_set.member?(s)
#     end
#     intersects
#   end
#
#   def difference(other_set)
#     set.each do |s|
#       set.delete(s) if other_set.member?(s)
#     end
#   end
#
#   def union(other_set)
#     set.each do |s|
#       other_set.add(s) unless other_set.member?(s)
#     end
#
#     other_set
#   end
#
#   def ==(other_set)
#     return false if set.length != other_set.length
#     set.each do |s|
#       if other_set.member?(s)
#         other_set.delete(s)
#       end
#     end
#
#     other_set.length == 0 ? true : false
#   end
#
#   def add(num)
#     return set if set.member?(num)
#     set << num
#   end
#
#   def delete(num)
#     set.filter! { |s| s != num }
#   end
#
#   def length
#     set.length
#   end
#
#   def each(&block)
#     set.each(&block)
#   end
# end