# This is a custom exception that you can use in your code
class NotMovieClubMemberError < RuntimeError
end

class Moviegoer
  def initialize(age, member: false)
    @age = age
    @member = member
  end

  def ticket_price
    discount_age_limit = 60
    @age >= discount_age_limit ? 10 : 15
  end

  def watch_scary_movie?
    adult_age_limit = 18
    @age >= adult_age_limit
  end

  # Popcorn is 🍿
  def claim_free_popcorn!
    @member ? '🍿' : (raise NotMovieClubMemberError)
  end
end
