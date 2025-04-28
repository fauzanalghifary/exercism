class Garden
  PLANT_MAPPING = {
    "G": "grass",
    "C": "clover",
    "R": "radishes",
    "V": "violets"
  }

  STUDENTS = %w(alice bob charlie david eve fred ginny
    harriet ileana joseph kincaid larry)

  # @param[String]garden
  def initialize(garden, students = STUDENTS)
    @diagram = garden.split"\n"
    @students = students.map(&:downcase).sort
  end

  def method_missing(name, *args)
    results = []
    idx = @students.index(name.to_s)
    @diagram.each do |row|
      results << PLANT_MAPPING[row[idx*2].to_sym].to_sym
      results << PLANT_MAPPING[row[idx*2+1].to_sym].to_sym
    end
    results
  end
end