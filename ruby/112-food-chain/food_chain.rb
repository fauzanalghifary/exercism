class FoodChain
  ANIMALS = [
    { name: 'fly', verse: nil },
    { name: 'spider', verse: 'It wriggled and jiggled and tickled inside her.' },
    { name: 'bird', verse: 'How absurd to swallow a bird!' },
    { name: 'cat', verse: 'Imagine that, to swallow a cat!' },
    { name: 'dog', verse: 'What a hog, to swallow a dog!' },
    { name: 'goat', verse: 'Just opened her throat and swallowed a goat!' },
    { name: 'cow', verse: "I don't know how she swallowed a cow!" },
    { name: 'horse', verse: "She's dead, of course!" }
  ].freeze

  def self.song
    verses = []

    ANIMALS.each_with_index do |animal, index|
      verse = "I know an old lady who swallowed a #{animal[:name]}.\n"

      if animal[:name] == 'horse'
        verse += animal[:verse]
      else
        verse += "#{animal[:verse]}\n" if animal[:verse]

        # Build the chain of animals
        (index).downto(1) do |i|
          caught = ANIMALS[i-1][:name]
          if caught == 'spider'
            verse += "She swallowed the #{ANIMALS[i][:name]} to catch the #{caught} that wriggled and jiggled and tickled inside her.\n"
          else
            verse += "She swallowed the #{ANIMALS[i][:name]} to catch the #{caught}.\n"
          end
        end

        verse += "I don't know why she swallowed the fly. Perhaps she'll die."
      end

      verses << verse
    end

    verses.join("\n\n") + "\n"
  end
end
