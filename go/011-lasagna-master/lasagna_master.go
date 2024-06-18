package lasagna

func PreparationTime(layers []string, avgTimePerLayer int) int {
	if avgTimePerLayer == 0 {
		return len(layers) * 2
	}
	return len(layers) * avgTimePerLayer
}

func Quantities(layers []string) (int, float64) {
	var noodlesAmount int
	var sauceAmount int
	for _, layer := range layers {
		if layer == "noodles" {
			noodlesAmount++
		} else if layer == "sauce" {
			sauceAmount++
		}
	}
	return noodlesAmount * 50, float64(sauceAmount) * 0.2
}

func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(quantities []float64, portionAmount int) []float64 {
	multiplier := float64(portionAmount) / 2
	var finalQuantities []float64
	for _, q := range quantities {
		finalQuantities = append(finalQuantities, q*multiplier)
	}
	return finalQuantities
}
