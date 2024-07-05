class Tournament

  @header = <<~HEADER
    Team                           | MP |  W |  D |  L |  P
  HEADER

  def self.calculate_point(result, reverse: false)
    return 3 if (result == 'win' && !reverse) || (result == 'loss' && reverse)
    return 1 if result == 'draw'

    0 if (result == 'loss' && !reverse) || (result == 'win' && reverse)
  end

  def self.tally(input)
    input_data = input.split("\n").map { |item| item.split(';') }
    teams = {}
    input_data.each do |data|
      team_a_name = data[0]
      team_b_name = data[1]
      result = data[2]

      teams[team_a_name] ||= {
        name: team_a_name,
        mp: 0,
        w: 0,
        d: 0,
        l: 0,
        p: 0
      }

      team_a = teams[team_a_name]
      team_a[:mp] += 1
      team_a[:w] += 1 if result == 'win'
      team_a[:d] += 1 if result == 'draw'
      team_a[:l] += 1 if result == 'loss'
      team_a[:p] += self.calculate_point(result)

      teams[team_b_name] ||= {
        name: team_b_name,
        mp: 0,
        w: 0,
        d: 0,
        l: 0,
        p: 0
      }

      team_b = teams[team_b_name]
      team_b[:mp] += 1
      team_b[:w] += 1 if result == 'loss'
      team_b[:d] += 1 if result == 'draw'
      team_b[:l] += 1 if result == 'win'
      team_b[:p] += self.calculate_point(result, reverse: true)
    end

    sorted_teams = teams.values.sort_by { |team| [-team[:p], team[:name]] }
    result = @header
    sorted_teams.each do |team|
      table = <<~TABLE
        #{team[:name].ljust(30)} | #{team[:mp].to_s.rjust(2)} | #{team[:w].to_s.rjust(2)} | #{team[:d].to_s.rjust(2)} | #{team[:l].to_s.rjust(2)} | #{team[:p].to_s.rjust(2)}
      TABLE
      result += table
    end

    result
  end
end