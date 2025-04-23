class School
  def initialize
    @students = {}
  end

  def roster
    @students.flat_map do |grade, names|
      names.sort.map { |name| [grade.to_i, name] }
    end.sort_by { |grade, name| [grade, name] }.map { |_, name| name }
  end

  # @param[String]name
  # @param[Numeric]grade
  def add(name, grade)
    return false if roster.include?(name)
    if @students[grade]
      return false if @students[grade].include?(name)
      @students[grade] << name
    else
      @students[grade] = [name]
    end

    true
  end

  def grade(grade_num)
    return [] unless @students[grade_num]
    @students.fetch(grade_num).sort
  end
end