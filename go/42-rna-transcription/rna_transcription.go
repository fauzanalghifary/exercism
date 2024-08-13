package strand

var DNAtoRNA = map[rune]string{
	'G': "C",
	'C': "G",
	'T': "A",
	'A': "U",
}

func ToRNA(dna string) string {
	result := ""

	for _, nucleotide := range dna {
		result += DNAtoRNA[nucleotide]
	}

	return result
}

// USING STRINGS.MAP
//func ToRNA(dna string) string {
//	return strings.Map(func(r rune) rune {
//		switch r {
//		case 'G':
//			return 'C'
//		case 'C':
//			return 'G'
//		case 'T':
//			return 'A'
//		case 'A':
//			return 'U'
//		}
//		return r
//	}, dna)
//}
