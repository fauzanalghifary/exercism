package zebra

type Solution struct {
	DrinksWater string
	OwnsZebra   string
}

const (
	red = iota
	green
	ivory
	yellow
	blue
)

const (
	englishman = iota
	spaniard
	ukrainian
	norwegian
	japanese
)

const (
	dog = iota
	snails
	fox
	horse
	zebra
)

const (
	coffee = iota
	tea
	milk
	orangeJuice
	water
)

const (
	oldGold = iota
	kools
	chesterfields
	luckyStrike
	parliaments
)

var nationalities = []string{"Englishman", "Spaniard", "Ukrainian", "Norwegian", "Japanese"}

func SolvePuzzle() Solution {
	// Generate all permutations and check constraints
	colors := []int{red, green, ivory, yellow, blue}
	nats := []int{englishman, spaniard, ukrainian, norwegian, japanese}
	pets := []int{dog, snails, fox, horse, zebra}
	drinks := []int{coffee, tea, milk, orangeJuice, water}
	smokes := []int{oldGold, kools, chesterfields, luckyStrike, parliaments}

	for _, colorPerm := range permutations(colors) {
		if !checkGreenRightOfIvory(colorPerm) {
			continue
		}

		for _, natPerm := range permutations(nats) {
			if !checkEnglishmanRed(colorPerm, natPerm) {
				continue
			}
			if natPerm[0] != norwegian {
				continue
			}
			if !checkNorwegianNextToBlue(colorPerm, natPerm) {
				continue
			}

			for _, petPerm := range permutations(pets) {
				if !checkSpaniardDog(natPerm, petPerm) {
					continue
				}

				for _, drinkPerm := range permutations(drinks) {
					if drinkPerm[2] != milk {
						continue
					}
					if !checkCoffeeGreen(colorPerm, drinkPerm) {
						continue
					}
					if !checkUkrainianTea(natPerm, drinkPerm) {
						continue
					}

					for _, smokePerm := range permutations(smokes) {
						if !checkOldGoldSnails(smokePerm, petPerm) {
							continue
						}
						if !checkKoolsYellow(smokePerm, colorPerm) {
							continue
						}
						if !checkLuckyStrikeOJ(smokePerm, drinkPerm) {
							continue
						}
						if !checkJapaneseParliaments(natPerm, smokePerm) {
							continue
						}
						if !checkChesterfieldsNextToFox(smokePerm, petPerm) {
							continue
						}
						if !checkKoolsNextToHorse(smokePerm, petPerm) {
							continue
						}

						// Found solution
						var result Solution
						for i := 0; i < 5; i++ {
							if drinkPerm[i] == water {
								result.DrinksWater = nationalities[natPerm[i]]
							}
							if petPerm[i] == zebra {
								result.OwnsZebra = nationalities[natPerm[i]]
							}
						}
						return result
					}
				}
			}
		}
	}

	return Solution{}
}

func permutations(arr []int) [][]int {
	var result [][]int
	permute(arr, 0, &result)
	return result
}

func permute(arr []int, start int, result *[][]int) {
	if start == len(arr)-1 {
		temp := make([]int, len(arr))
		copy(temp, arr)
		*result = append(*result, temp)
		return
	}

	for i := start; i < len(arr); i++ {
		arr[start], arr[i] = arr[i], arr[start]
		permute(arr, start+1, result)
		arr[start], arr[i] = arr[i], arr[start]
	}
}

func checkEnglishmanRed(colors, nats []int) bool {
	for i := 0; i < 5; i++ {
		if nats[i] == englishman && colors[i] != red {
			return false
		}
		if nats[i] != englishman && colors[i] == red {
			return false
		}
	}
	return true
}

func checkSpaniardDog(nats, pets []int) bool {
	for i := 0; i < 5; i++ {
		if nats[i] == spaniard && pets[i] != dog {
			return false
		}
		if nats[i] != spaniard && pets[i] == dog {
			return false
		}
	}
	return true
}

func checkCoffeeGreen(colors, drinks []int) bool {
	for i := 0; i < 5; i++ {
		if colors[i] == green && drinks[i] != coffee {
			return false
		}
		if colors[i] != green && drinks[i] == coffee {
			return false
		}
	}
	return true
}

func checkUkrainianTea(nats, drinks []int) bool {
	for i := 0; i < 5; i++ {
		if nats[i] == ukrainian && drinks[i] != tea {
			return false
		}
		if nats[i] != ukrainian && drinks[i] == tea {
			return false
		}
	}
	return true
}

func checkGreenRightOfIvory(colors []int) bool {
	for i := 0; i < 5; i++ {
		if colors[i] == ivory {
			if i == 4 || colors[i+1] != green {
				return false
			}
		}
	}
	return true
}

func checkOldGoldSnails(smokes, pets []int) bool {
	for i := 0; i < 5; i++ {
		if smokes[i] == oldGold && pets[i] != snails {
			return false
		}
		if smokes[i] != oldGold && pets[i] == snails {
			return false
		}
	}
	return true
}

func checkKoolsYellow(smokes, colors []int) bool {
	for i := 0; i < 5; i++ {
		if smokes[i] == kools && colors[i] != yellow {
			return false
		}
		if smokes[i] != kools && colors[i] == yellow {
			return false
		}
	}
	return true
}

func checkLuckyStrikeOJ(smokes, drinks []int) bool {
	for i := 0; i < 5; i++ {
		if smokes[i] == luckyStrike && drinks[i] != orangeJuice {
			return false
		}
		if smokes[i] != luckyStrike && drinks[i] == orangeJuice {
			return false
		}
	}
	return true
}

func checkJapaneseParliaments(nats, smokes []int) bool {
	for i := 0; i < 5; i++ {
		if nats[i] == japanese && smokes[i] != parliaments {
			return false
		}
		if nats[i] != japanese && smokes[i] == parliaments {
			return false
		}
	}
	return true
}

func checkNorwegianNextToBlue(colors, nats []int) bool {
	for i := 0; i < 5; i++ {
		if nats[i] == norwegian {
			if i == 0 && colors[1] != blue {
				return false
			}
			if i == 4 && colors[3] != blue {
				return false
			}
			if i > 0 && i < 4 && colors[i-1] != blue && colors[i+1] != blue {
				return false
			}
		}
	}
	return true
}

func checkChesterfieldsNextToFox(smokes, pets []int) bool {
	for i := 0; i < 5; i++ {
		if smokes[i] == chesterfields {
			hasNeighborFox := false
			if i > 0 && pets[i-1] == fox {
				hasNeighborFox = true
			}
			if i < 4 && pets[i+1] == fox {
				hasNeighborFox = true
			}
			if !hasNeighborFox {
				return false
			}
		}
	}
	return true
}

func checkKoolsNextToHorse(smokes, pets []int) bool {
	for i := 0; i < 5; i++ {
		if smokes[i] == kools {
			hasNeighborHorse := false
			if i > 0 && pets[i-1] == horse {
				hasNeighborHorse = true
			}
			if i < 4 && pets[i+1] == horse {
				hasNeighborHorse = true
			}
			if !hasNeighborHorse {
				return false
			}
		}
	}
	return true
}
