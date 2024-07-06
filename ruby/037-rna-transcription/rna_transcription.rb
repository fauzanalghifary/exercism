class Complement
  DNA_TO_RNA = {
    G: 'C',
    C: 'G',
    T: 'A',
    A: 'U'
  }.freeze

  def self.of_dna(dna)
    # rna = ''
    # dna.each_char do |char|
    #   rna += DNA_TO_RNA[char.to_sym] || ''
    # end
    # rna
    dna.tr('GCTA', 'CGAU')
  end
end