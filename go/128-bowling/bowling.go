package bowling

import "errors"

// Game represents a bowling game
type Game struct {
	rolls []int
}

func NewGame() *Game {
	return &Game{
		rolls: make([]int, 0),
	}
}

func (g *Game) Roll(pins int) error {
	if pins < 0 {
		return errors.New("negative roll is invalid")
	}

	if pins > 10 {
		return errors.New("pin count exceeds pins on the lane")
	}

	if g.isGameOver() {
		return errors.New("cannot roll after game is over")
	}

	// Validate based on current game state
	rollIndex := len(g.rolls)

	// Before frame 10 (rolls 0-17 max)
	if rollIndex < 18 {
		// Check if this is the second roll of a frame
		if rollIndex%2 == 1 {
			// Get the first roll of this frame
			firstRoll := g.rolls[rollIndex-1]
			// If first roll wasn't a strike, check total doesn't exceed 10
			if firstRoll < 10 && firstRoll+pins > 10 {
				return errors.New("pin count exceeds pins on the lane")
			}
		}
	} else {
		// Frame 10 (rolls 18+)
		if rollIndex == 18 {
			// First roll of frame 10 - no special validation needed
		} else if rollIndex == 19 {
			// Second roll of frame 10
			firstRoll := g.rolls[18]
			if firstRoll < 10 && firstRoll+pins > 10 {
				return errors.New("pin count exceeds pins on the lane")
			}
		} else if rollIndex == 20 {
			// Third roll (bonus roll) of frame 10
			firstRoll := g.rolls[18]
			secondRoll := g.rolls[19]

			// If first roll was a strike, second roll determines validation
			if firstRoll == 10 {
				// If second roll wasn't a strike, third roll can't make total > 10
				if secondRoll < 10 && secondRoll+pins > 10 {
					return errors.New("pin count exceeds pins on the lane")
				}
			}
			// If first two rolls were a spare, bonus roll is allowed
		}
	}

	g.rolls = append(g.rolls, pins)
	return nil
}

func (g *Game) isGameOver() bool {
	rollCount := len(g.rolls)

	// Less than 18 rolls means we haven't finished 9 frames
	if rollCount < 18 {
		return false
	}

	// Exactly 18 rolls - need to check if we're in frame 10
	if rollCount == 18 {
		return false
	}

	// In frame 10 (rolls 18+)
	if rollCount == 19 {
		// If first roll was a strike, we need 2 more rolls
		if g.rolls[18] == 10 {
			return false
		}
		// If no strike, check if we have a spare
		// If spare, need one more roll
		return false
	}

	if rollCount == 20 {
		firstRoll := g.rolls[18]
		secondRoll := g.rolls[19]

		// If first roll was a strike, we need 2 bonus rolls (total 3)
		if firstRoll == 10 {
			return false
		}

		// If we have a spare, we need 1 bonus roll (total 3)
		if firstRoll+secondRoll == 10 {
			return false
		}

		// No strike or spare in frame 10, game is over after 2 rolls
		return true
	}

	// 21 or more rolls
	return rollCount >= 21
}

func (g *Game) Score() (int, error) {
	if !g.isComplete() {
		return 0, errors.New("score cannot be taken until the end of the game")
	}

	score := 0
	rollIndex := 0

	for frame := 0; frame < 10; frame++ {
		if g.isStrike(rollIndex) {
			score += 10 + g.strikeBonus(rollIndex)
			rollIndex++
		} else if g.isSpare(rollIndex) {
			score += 10 + g.spareBonus(rollIndex)
			rollIndex += 2
		} else {
			score += g.rolls[rollIndex] + g.rolls[rollIndex+1]
			rollIndex += 2
		}
	}

	return score, nil
}

func (g *Game) isComplete() bool {
	rollCount := len(g.rolls)

	if rollCount < 10 {
		return false
	}

	// Count actual frames completed
	frameCount := 0
	rollIndex := 0

	for frameCount < 9 && rollIndex < rollCount {
		if g.rolls[rollIndex] == 10 {
			// Strike
			rollIndex++
		} else {
			// Normal frame (2 rolls)
			rollIndex += 2
		}
		frameCount++
	}

	// Now we should be at frame 10 (rollIndex should be at the start of frame 10)
	if frameCount < 9 {
		return false
	}

	// Check frame 10 completion
	if rollIndex >= rollCount {
		return false
	}

	firstRoll := g.rolls[rollIndex]

	// If strike in frame 10, need 2 more rolls (total 3 rolls in frame 10)
	if firstRoll == 10 {
		return rollCount > rollIndex+2
	}

	// Need at least 2 rolls in frame 10
	if rollIndex+1 >= rollCount {
		return false
	}

	secondRoll := g.rolls[rollIndex+1]

	// If spare in frame 10, need 1 more roll (total 3 rolls in frame 10)
	if firstRoll+secondRoll == 10 {
		return rollCount > rollIndex+2
	}

	// No strike or spare, just need 2 rolls
	return rollCount > rollIndex+1
}

func (g *Game) isStrike(rollIndex int) bool {
	return g.rolls[rollIndex] == 10
}

func (g *Game) isSpare(rollIndex int) bool {
	return g.rolls[rollIndex]+g.rolls[rollIndex+1] == 10
}

func (g *Game) strikeBonus(rollIndex int) int {
	return g.rolls[rollIndex+1] + g.rolls[rollIndex+2]
}

func (g *Game) spareBonus(rollIndex int) int {
	return g.rolls[rollIndex+2]
}
