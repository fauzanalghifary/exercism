package dndcharacter

import (
	"math"
	"math/rand"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return int(math.Floor(float64(score-10) / 2))
}

func Ability() int {
	return rand.Intn(15) + 3
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	char := Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: Ability(),
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
	}

	char.Hitpoints = 10 + Modifier(char.Constitution)

	return char
}

// Ability uses randomness to generate the score for an ability
//func Ability() int {
//	var result []int
//	for i := 0; i < 4; i++ {
//		result = append(result, rand.Intn(6)+1)
//	}
//
//	smallest := result[0]
//	for _, num := range result {
//		if num < smallest {
//			smallest = num
//		}
//	}
//	resultNum := 0
//	removed := false
//	for _, num := range result {
//		if num == smallest && !removed {
//			removed = true
//			continue
//		}
//		resultNum += num
//	}
//
//	return resultNum
//}
