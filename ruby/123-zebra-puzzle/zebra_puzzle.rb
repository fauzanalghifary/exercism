class ZebraPuzzle
  NATIONALITIES = %w[Norwegian Ukrainian Englishman Spaniard Japanese]
  COLORS = %w[red green ivory yellow blue]
  PETS = %w[dog snail fox horse zebra]
  DRINKS = ["coffee", "tea", "milk", "orange juice", "water"]
  HOBBIES = %w[dancing painting reading football chess]

  def self.water_drinker
    solution[:water_drinker]
  end

  def self.zebra_owner
    solution[:zebra_owner]
  end

  def self.solution
    @solution ||= solve
  end

  def self.solve
    # Try all permutations and find the one that satisfies all constraints
    NATIONALITIES.permutation.each do |nationalities|
      COLORS.permutation.each do |colors|
        next unless constraint_6(colors) # Green right of ivory
        next unless constraint_2(nationalities, colors) # Englishman in red
        next unless constraint_10(nationalities) # Norwegian in first house
        next unless constraint_15(nationalities, colors) # Norwegian next to blue

        PETS.permutation.each do |pets|
          next unless constraint_3(nationalities, pets) # Spaniard owns dog
          next unless constraint_7(pets) # Snail owner dances (checked with hobbies later)

          DRINKS.permutation.each do |drinks|
            next unless constraint_4(colors, drinks) # Green house drinks coffee
            next unless constraint_5(nationalities, drinks) # Ukrainian drinks tea
            next unless constraint_9(drinks) # Middle house drinks milk
            next unless constraint_13(drinks) # Football player drinks orange juice (checked with hobbies later)

            HOBBIES.permutation.each do |hobbies|
              next unless constraint_7_full(pets, hobbies) # Snail owner dances
              next unless constraint_8(colors, hobbies) # Yellow house painter
              next unless constraint_11(hobbies, pets) # Reader next to fox
              next unless constraint_12(hobbies, pets) # Painter next to horse
              next unless constraint_13_full(hobbies, drinks) # Football player drinks orange juice
              next unless constraint_14(nationalities, hobbies) # Japanese plays chess

              # Found the solution!
              water_idx = drinks.index("water")
              zebra_idx = pets.index("zebra")

              return {
                water_drinker: nationalities[water_idx],
                zebra_owner: nationalities[zebra_idx]
              }
            end
          end
        end
      end
    end

    nil
  end

  # Constraint 2: The Englishman lives in the red house
  def self.constraint_2(nationalities, colors)
    nationalities.index("Englishman") == colors.index("red")
  end

  # Constraint 3: The Spaniard owns the dog
  def self.constraint_3(nationalities, pets)
    nationalities.index("Spaniard") == pets.index("dog")
  end

  # Constraint 4: The person in the green house drinks coffee
  def self.constraint_4(colors, drinks)
    colors.index("green") == drinks.index("coffee")
  end

  # Constraint 5: The Ukrainian drinks tea
  def self.constraint_5(nationalities, drinks)
    nationalities.index("Ukrainian") == drinks.index("tea")
  end

  # Constraint 6: The green house is immediately to the right of the ivory house
  def self.constraint_6(colors)
    ivory_idx = colors.index("ivory")
    green_idx = colors.index("green")
    green_idx == ivory_idx + 1
  end

  # Constraint 7: The snail owner likes to go dancing (partial check)
  def self.constraint_7(pets)
    true # Will be fully checked with hobbies
  end

  # Constraint 7 full: The snail owner likes to go dancing
  def self.constraint_7_full(pets, hobbies)
    pets.index("snail") == hobbies.index("dancing")
  end

  # Constraint 8: The person in the yellow house is a painter
  def self.constraint_8(colors, hobbies)
    colors.index("yellow") == hobbies.index("painting")
  end

  # Constraint 9: The person in the middle house drinks milk
  def self.constraint_9(drinks)
    drinks[2] == "milk"
  end

  # Constraint 10: The Norwegian lives in the first house
  def self.constraint_10(nationalities)
    nationalities[0] == "Norwegian"
  end

  # Constraint 11: The person who enjoys reading lives in the house next to the person with the fox
  def self.constraint_11(hobbies, pets)
    reader_idx = hobbies.index("reading")
    fox_idx = pets.index("fox")
    (reader_idx - fox_idx).abs == 1
  end

  # Constraint 12: The painter's house is next to the house with the horse
  def self.constraint_12(hobbies, pets)
    painter_idx = hobbies.index("painting")
    horse_idx = pets.index("horse")
    (painter_idx - horse_idx).abs == 1
  end

  # Constraint 13: The person who plays football drinks orange juice (partial)
  def self.constraint_13(drinks)
    true # Will be fully checked with hobbies
  end

  # Constraint 13 full: The person who plays football drinks orange juice
  def self.constraint_13_full(hobbies, drinks)
    hobbies.index("football") == drinks.index("orange juice")
  end

  # Constraint 14: The Japanese person plays chess
  def self.constraint_14(nationalities, hobbies)
    nationalities.index("Japanese") == hobbies.index("chess")
  end

  # Constraint 15: The Norwegian lives next to the blue house
  def self.constraint_15(nationalities, colors)
    norwegian_idx = nationalities.index("Norwegian")
    blue_idx = colors.index("blue")
    (norwegian_idx - blue_idx).abs == 1
  end
end
