package twobucket

import (
	"errors"
)

type state struct {
	bucket1, bucket2, moves int
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Solve(sizeBucketOne, sizeBucketTwo, goalAmount int, startBucket string) (
	string,
	int,
	int,
	error,
) {
	if sizeBucketOne <= 0 {
		return "", 0, 0, errors.New("invalid first bucket size")
	}
	if sizeBucketTwo <= 0 {
		return "", 0, 0, errors.New("invalid second bucket size")
	}
	if goalAmount <= 0 {
		return "", 0, 0, errors.New("invalid goal amount")
	}
	if startBucket != "one" && startBucket != "two" {
		return "", 0, 0, errors.New("invalid start bucket name")
	}

	if goalAmount > sizeBucketOne && goalAmount > sizeBucketTwo {
		return "", 0, 0, errors.New("impossible")
	}
	if goalAmount%gcd(sizeBucketOne, sizeBucketTwo) != 0 {
		return "", 0, 0, errors.New("impossible")
	}

	// BFS to find the shortest path
	visited := make(map[state]bool)
	queue := []state{}

	// Initialize based on the start bucket
	if startBucket == "one" {
		queue = append(queue, state{sizeBucketOne, 0, 1})
	} else {
		queue = append(queue, state{0, sizeBucketTwo, 1})
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// Skip if already visited
		if visited[state{current.bucket1, current.bucket2, 0}] {
			continue
		}
		visited[state{current.bucket1, current.bucket2, 0}] = true

		// Check if the goal is reached
		if current.bucket1 == goalAmount {
			return "one", current.moves, current.bucket2, nil
		}
		if current.bucket2 == goalAmount {
			return "two", current.moves, current.bucket1, nil
		}

		// Constraint: cannot have a start bucket empty and another bucket full after the first move
		cannotReachState := func(b1, b2 int) bool {
			if startBucket == "one" && b1 == 0 && b2 == sizeBucketTwo {
				return true
			}
			if startBucket == "two" && b2 == 0 && b1 == sizeBucketOne {
				return true
			}
			return false
		}

		// Generate next states
		nextStates := []state{}

		// Fill bucket 1
		if current.bucket1 < sizeBucketOne && !cannotReachState(sizeBucketOne, current.bucket2) {
			nextStates = append(
				nextStates,
				state{sizeBucketOne, current.bucket2, current.moves + 1},
			)
		}

		// Fill bucket 2
		if current.bucket2 < sizeBucketTwo && !cannotReachState(current.bucket1, sizeBucketTwo) {
			nextStates = append(
				nextStates,
				state{current.bucket1, sizeBucketTwo, current.moves + 1},
			)
		}

		// Empty bucket 1
		if current.bucket1 > 0 && !cannotReachState(0, current.bucket2) {
			nextStates = append(nextStates, state{0, current.bucket2, current.moves + 1})
		}

		// Empty bucket 2
		if current.bucket2 > 0 && !cannotReachState(current.bucket1, 0) {
			nextStates = append(nextStates, state{current.bucket1, 0, current.moves + 1})
		}

		// Pour from bucket 1 to bucket 2
		if current.bucket1 > 0 && current.bucket2 < sizeBucketTwo {
			pour := min(current.bucket1, sizeBucketTwo-current.bucket2)
			newB1 := current.bucket1 - pour
			newB2 := current.bucket2 + pour
			if !cannotReachState(newB1, newB2) {
				nextStates = append(nextStates, state{newB1, newB2, current.moves + 1})
			}
		}

		// Pour from bucket 2 to bucket 1
		if current.bucket2 > 0 && current.bucket1 < sizeBucketOne {
			pour := min(current.bucket2, sizeBucketOne-current.bucket1)
			newB1 := current.bucket1 + pour
			newB2 := current.bucket2 - pour
			if !cannotReachState(newB1, newB2) {
				nextStates = append(nextStates, state{newB1, newB2, current.moves + 1})
			}
		}

		queue = append(queue, nextStates...)
	}

	return "", 0, 0, errors.New("impossible")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
