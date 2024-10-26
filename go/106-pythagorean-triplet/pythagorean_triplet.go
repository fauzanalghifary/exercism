package pythagorean

type Triplet [3]int

// Range generates list of all Pythagorean triplets with side lengths
// in the provided range.
func Range(min, max int) []Triplet {
	var triplets []Triplet
	for a := min; a <= max; a++ {
		for b := a + 1; b <= max; b++ {
			for c := b + 1; c <= max; c++ {
				if a*a+b*b == c*c {
					triplets = append(triplets, Triplet{a, b, c})
				}
			}
		}
	}
	return triplets
}

// Sum returns a list of all Pythagorean triplets with a certain perimeter.
func Sum(p int) []Triplet {
	candidates := Range(1, p/2)

	var triplets []Triplet
	for _, triplet := range candidates {
		if triplet[0]+triplet[1]+triplet[2] == p {
			triplets = append(triplets, triplet)
		}
	}
	return triplets
}
