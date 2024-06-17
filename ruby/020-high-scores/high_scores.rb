class HighScores
  attr_reader :scores

  def initialize(scores)
    @scores = scores
  end

  def latest
    # scores[-1]
    scores.last
  end

  def personal_best
    scores.max
  end

  def personal_top_three
    # ordered_scores = scores.sort { |a, b| b <=> a }
    # ordered_scores.first(3)
    scores.max(3)
  end

  def latest_is_personal_best?
    # the_best = personal_best
    # the_latest = latest
    # the_best == the_latest
    latest == personal_best
  end
end