package protein

import (
	"errors"
)

var ErrStop = errors.New("stop codon")
var ErrInvalidBase = errors.New("invalid base")

var codonsToProtein = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

const codonLength = 3

func FromRNA(rna string) ([]string, error) {
	var result []string
	for i := 0; i < len(rna); i += codonLength {
		protein, err := FromCodon(rna[i : i+codonLength])

		if err != nil {
			if errors.Is(err, ErrStop) {
				break
			}

			if errors.Is(err, ErrInvalidBase) {
				return result, ErrInvalidBase
			}
		}

		result = append(result, protein)
	}

	return result, nil
}

func FromCodon(codon string) (string, error) {
	protein := codonsToProtein[codon]

	if protein == "" {
		return "", ErrInvalidBase
	}

	if protein == "STOP" {
		return "", ErrStop
	}

	return protein, nil
}
