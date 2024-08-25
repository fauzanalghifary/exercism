package triangle

type Kind int

const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

func KindFromSides(a, b, c float64) Kind {
	var k Kind

	if a == 0 || b == 0 || c == 0 {
		k = NaT
	} else if a+b < c || a+c < b || b+c < a {
		k = NaT
	} else if a == b && b == c {
		k = Equ
	} else if a != b && a != c && b != c {
		k = Sca
	} else {
		k = Iso
	}

	return k
}
