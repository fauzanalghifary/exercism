class Microwave
  def initialize(time)
    @time = time
  end

  def timer
    minutes = @time / 100 # get everything except the last 2 digits
    seconds = @time % 100 # get last 2 digits

    minutes += seconds / 60
    seconds = seconds % 60

    format('%02d:%02d', minutes, seconds)
  end
end
