package space

type Planet string

func Age(seconds float64, planet Planet) float64 {
	yearOnEarth := seconds / 31557600

	if planet == "Earth" {
		return yearOnEarth
	} else if planet == "Mercury" {
		return yearOnEarth / 0.2408467
	} else if planet == "Venus" {
		return yearOnEarth / 0.61519726
	} else if planet == "Mars" {
		return yearOnEarth / 1.8808158
	} else if planet == "Jupiter" {
		return yearOnEarth / 11.862615
	} else if planet == "Saturn" {
		return yearOnEarth / 29.447498
	} else if planet == "Uranus" {
		return yearOnEarth / 84.016846
	} else if planet == "Neptune" {
		return yearOnEarth / 164.79132
	}

	return -1.000000
}
