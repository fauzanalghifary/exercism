package cipher

import "strings"

// Define the shift and vigenere types here.
// Both types should satisfy the Cipher interface.
type shift struct {
	distance int
}

type vigenere struct {
	distances []int
}

func NewCaesar() Cipher {
	return &shift{distance: 3}
}

func NewShift(distance int) Cipher {
	if distance < -25 || distance == 0 || distance > 25 {
		return nil
	}
	return &shift{distance: distance}
}

func (c shift) Encode(input string) string {
	var result strings.Builder
	result.Grow(len(input))

	for _, r := range strings.ToLower(input) {
		switch {
		case r >= 'a' && r <= 'z':
			result.WriteRune(shiftRune(r, c.distance))
		}
	}
	return result.String()
}

func (c shift) Decode(input string) string {
	var result strings.Builder
	result.Grow(len(input))

	for _, r := range strings.ToLower(input) {
		switch {
		case r >= 'a' && r <= 'z':
			result.WriteRune(shiftRune(r, -c.distance))
		}
	}
	return result.String()
}

func NewVigenere(key string) Cipher {
	if key == "" {
		return nil
	}
	var valid bool = false
	distances := make([]int, 0, len(key))
	for _, r := range key {
		if r < 'a' || r > 'z' {
			valid = false
			break
		}
		if r != 'a' {
			valid = true
		}
		distances = append(distances, int(r-'a'))
	}
	if !valid {
		return nil
	}
	return &vigenere{distances: distances}
}

func (v vigenere) Encode(input string) string {
	var result strings.Builder
	result.Grow(len(input))

	var idx int
	for _, r := range strings.ToLower(input) {
		switch {
		case r >= 'a' && r <= 'z':
			result.WriteRune(shiftRune(r, v.distances[idx%len(v.distances)]))
			idx++
		}
	}
	return result.String()
}

func (v vigenere) Decode(input string) string {
	var result strings.Builder
	result.Grow(len(input))
	var idx int
	for _, r := range strings.ToLower(input) {
		switch {
		case r >= 'a' && r <= 'z':
			result.WriteRune(shiftRune(r, -v.distances[idx%len(v.distances)]))
			idx++
		}
	}
	return result.String()
}

func shiftRune(r rune, distance int) rune {
	if distance < 0 {
		distance = 26 + distance
	}
	idx := (int(r-'a') + distance) % 26
	return rune('a' + idx)
}
