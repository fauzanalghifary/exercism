package ledger

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"strings"
	"time"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

var translations = map[string]map[string]string{
	"nl-NL": {
		"Date":                  "Datum",
		"Description":           "Omschrijving",
		"Change":                "Verandering",
		"date_format":           "02-01-2006",
		"price_format":          "%s %s ",
		"price_format_negative": "%s %s-",
		"thou":                  ".",
		"dec":                   ",",
	},
	"en-US": {
		"Date":                  "Date",
		"Description":           "Description",
		"Change":                "Change",
		"date_format":           "01/02/2006",
		"price_format":          "%s%s ",
		"price_format_negative": "(%s%s)",
		"thou":                  ",",
		"dec":                   ".",
	},
}
var curr = map[string]string{
	"EUR": "€",
	"USD": "$",
}

func formatDate(d, locale string) (string, bool) {
	date, err := time.Parse("2006-01-02", d)
	if err != nil {
		return "", false
	}
	return date.Format(translations[locale]["date_format"]), true
}

func formatDesc(de string) string {
	runes := []rune(de)
	if len(runes) > 25 {
		return string(runes[:22]) + "..."
	}
	return fmt.Sprintf("%-25s", de)
}

func formatValue(currency string, value int, locale string) string {
	formatKey := "price_format"
	if value < 0 {
		formatKey += "_negative"
	}
	v := math.Abs(float64(value) / 100.0)
	vf := fmt.Sprintf("%.2f", v)
	parts := strings.Split(vf, ".")
	if len(parts[0]) > 3 {
		s := len(parts[0]) - 3
		parts[0] = parts[0][:s] + translations[locale]["thou"] + parts[0][s:]
	}
	vf = strings.Join(parts, translations[locale]["dec"])
	return fmt.Sprintf(translations[locale][formatKey], curr[currency], vf)
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	if currency != "USD" && currency != "EUR" {
		return "", errors.New("invalid currency")
	}

	if locale != "en-US" && locale != "nl-NL" {
		return "", errors.New("invalid locale")
	}

	entriesCopy := append([]Entry{}, entries...)
	sort.Slice(
		entriesCopy, func(i, j int) bool {
			if entriesCopy[i].Date != entriesCopy[j].Date {
				return entriesCopy[i].Date < entriesCopy[j].Date
			}
			if entriesCopy[i].Description != entriesCopy[j].Description {
				return entriesCopy[i].Description < entriesCopy[j].Description
			}
			return entriesCopy[i].Change < entriesCopy[j].Change
		},
	)

	s := fmt.Sprintf(
		"%-10s | %-25s | %s\n",
		translations[locale]["Date"],
		translations[locale]["Description"],
		translations[locale]["Change"],
	)

	for _, entry := range entriesCopy {
		d, ok := formatDate(entry.Date, locale)
		if !ok {
			return "", errors.New("")
		}
		de := formatDesc(entry.Description)
		a := formatValue(currency, entry.Change, locale)
		s += fmt.Sprintf("%-10s | %-25s | %13s\n", d, de, a)
	}

	return s, nil
}
