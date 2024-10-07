package allergies

import "slices"

var scoreToAllergen = map[uint]string{
	1:   "eggs",
	2:   "peanuts",
	4:   "shellfish",
	8:   "strawberries",
	16:  "tomatoes",
	32:  "chocolate",
	64:  "pollen",
	128: "cats",
}

func Allergies(allergies uint) []string {
	var allergiesList []string
	if allergies == 0 {
		return allergiesList
	}

	allergiesScore := allergies % 256

	for i := 128; i > 0; i = i / 2 {
		isAllergy := int(allergiesScore) / i
		if isAllergy >= 1 {
			allergiesList = append(allergiesList, scoreToAllergen[uint(i)])
			allergiesScore -= uint(i)
		}

	}

	return allergiesList
}

func AllergicTo(allergies uint, allergen string) bool {
	allergiesList := Allergies(allergies)
	return slices.Contains(allergiesList, allergen)
}
