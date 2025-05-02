class House
  PIECES = [
    ['the malt', 'that lay in'],
    ['the rat', 'that ate'],
    ['the cat', 'that killed'],
    ['the dog', 'that worried'],
    ['the cow with the crumpled horn', 'that tossed'],
    ['the maiden all forlorn', 'that milked'],
    ['the man all tattered and torn', 'that kissed'],
    ['the priest all shaven and shorn', 'that married'],
    ['the rooster that crowed in the morn', 'that woke'],
    ['the farmer sowing his corn', 'that kept'],
    ['the horse and the hound and the horn', 'that belonged to']
  ].freeze

  def self.recite
    verses = ["This is the house that Jack built."]

    (0...PIECES.length).each do |i|
      current_verse_lines = ["This is #{PIECES[i][0]}"]
      (i.downto(0)).each do |j|
        line = if j == 0
                 "#{PIECES[j][1]} the house that Jack built."
               else
                 "#{PIECES[j][1]} #{PIECES[j - 1][0]}"
               end
        current_verse_lines << line
      end
      verses << current_verse_lines.join("\n")
    end

    verses.join("\n\n") + "\n"
  end
end