package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

type Robot struct {
	name string
}

var (
	usedNames = make(map[string]struct{})
	rng       = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func generateName() string {
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits := "0123456789"
	return fmt.Sprintf(
		"%c%c%c%c%c",
		letters[rng.Intn(len(letters))],
		letters[rng.Intn(len(letters))],
		digits[rng.Intn(len(digits))],
		digits[rng.Intn(len(digits))],
		digits[rng.Intn(len(digits))],
	)
}

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		for {
			name := generateName()
			if _, exists := usedNames[name]; !exists {
				r.name = name
				usedNames[name] = struct{}{}
				break
			}
		}
	}
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}
