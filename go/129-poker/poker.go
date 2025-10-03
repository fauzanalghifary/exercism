package poker

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

// Card represents a poker card
type Card struct {
	rank int
	suit rune
}

// Hand represents a poker hand
type Hand struct {
	cards    []Card
	original string
}

// HandRank represents the ranking of a hand
type HandRank int

const (
	HighCard HandRank = iota
	OnePair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
)

var rankMap = map[string]int{
	"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "10": 10,
	"J": 11, "Q": 12, "K": 13, "A": 14,
}

var validSuits = map[string]bool{
	"♤": true, "♡": true, "♢": true, "♧": true,
}

func parseCard(cardStr string) (Card, error) {
	if len(cardStr) < 2 {
		return Card{}, errors.New("invalid card format")
	}

	// Find where the suit starts (multi-byte Unicode character)
	// Suits are 3-byte UTF-8 characters
	suitStart := -1
	for i := len(cardStr) - 1; i >= 0; i-- {
		if (cardStr[i] & 0xC0) != 0x80 { // Start of a UTF-8 character
			suitStart = i
			break
		}
	}

	if suitStart == -1 || suitStart == 0 {
		return Card{}, errors.New("invalid card format")
	}

	suitStr := cardStr[suitStart:]
	if !validSuits[suitStr] {
		return Card{}, errors.New("invalid suit")
	}

	// Extract rank (everything before suit)
	rankStr := cardStr[:suitStart]
	if rankStr == "" {
		return Card{}, errors.New("missing rank")
	}

	rank, ok := rankMap[rankStr]
	if !ok {
		return Card{}, errors.New("invalid rank")
	}

	// Store first rune of suit
	suit := []rune(suitStr)[0]
	return Card{rank: rank, suit: suit}, nil
}

func parseHand(handStr string) (Hand, error) {
	cardStrs := strings.Fields(handStr)
	if len(cardStrs) != 5 {
		return Hand{}, errors.New("hand must have exactly 5 cards")
	}

	cards := make([]Card, 5)
	for i, cardStr := range cardStrs {
		card, err := parseCard(cardStr)
		if err != nil {
			return Hand{}, err
		}
		cards[i] = card
	}

	return Hand{cards: cards, original: handStr}, nil
}

func (h *Hand) getRankCounts() []int {
	counts := make(map[int]int)
	for _, card := range h.cards {
		counts[card.rank]++
	}

	result := make([]int, 0, len(counts))
	for _, count := range counts {
		result = append(result, count)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(result)))
	return result
}

func (h *Hand) isFlush() bool {
	suit := h.cards[0].suit
	for _, card := range h.cards[1:] {
		if card.suit != suit {
			return false
		}
	}
	return true
}

func (h *Hand) isStraight() bool {
	ranks := make([]int, len(h.cards))
	for i, card := range h.cards {
		ranks[i] = card.rank
	}
	sort.Ints(ranks)

	// Check for A-2-3-4-5 straight (wheel) first
	if ranks[0] == 2 && ranks[1] == 3 && ranks[2] == 4 && ranks[3] == 5 && ranks[4] == 14 {
		return true
	}

	// Check for regular straight
	for i := 1; i < len(ranks); i++ {
		if ranks[i] != ranks[i-1]+1 {
			return false
		}
	}
	return true
}

func (h *Hand) getHandRank() HandRank {
	counts := h.getRankCounts()
	isFlush := h.isFlush()
	isStraight := h.isStraight()

	if isStraight && isFlush {
		return StraightFlush
	}
	if len(counts) == 2 && counts[0] == 4 {
		return FourOfAKind
	}
	if len(counts) == 2 && counts[0] == 3 {
		return FullHouse
	}
	if isFlush {
		return Flush
	}
	if isStraight {
		return Straight
	}
	if len(counts) == 3 && counts[0] == 3 {
		return ThreeOfAKind
	}
	if len(counts) == 3 && counts[0] == 2 && counts[1] == 2 {
		return TwoPair
	}
	if len(counts) == 4 && counts[0] == 2 {
		return OnePair
	}
	return HighCard
}

func (h *Hand) getTieBreakers() []int {
	rankCounts := make(map[int]int)
	for _, card := range h.cards {
		rankCounts[card.rank]++
	}

	// Group ranks by count
	type rankCount struct {
		rank  int
		count int
	}
	var groups []rankCount
	for rank, count := range rankCounts {
		groups = append(groups, rankCount{rank: rank, count: count})
	}

	// Sort by count (descending) then by rank (descending)
	sort.Slice(
		groups, func(i, j int) bool {
			if groups[i].count != groups[j].count {
				return groups[i].count > groups[j].count
			}
			return groups[i].rank > groups[j].rank
		},
	)

	// Build tie breaker array
	tieBreakers := make([]int, 0, 5)
	for _, g := range groups {
		tieBreakers = append(tieBreakers, g.rank)
	}

	// Special case for A-2-3-4-5 straight (wheel)
	if h.isStraight() {
		ranks := make([]int, len(h.cards))
		for i, card := range h.cards {
			ranks[i] = card.rank
		}
		sort.Ints(ranks)
		if ranks[0] == 2 && ranks[4] == 14 {
			// In A-2-3-4-5 straight, Ace is low, so use 5 as the high card
			return []int{5}
		}
		return []int{ranks[4]}
	}

	return tieBreakers
}

func compareHands(h1, h2 *Hand) int {
	rank1 := h1.getHandRank()
	rank2 := h2.getHandRank()

	if rank1 != rank2 {
		if rank1 > rank2 {
			return 1
		}
		return -1
	}

	// Same rank, compare tie breakers
	tb1 := h1.getTieBreakers()
	tb2 := h2.getTieBreakers()

	for i := 0; i < len(tb1) && i < len(tb2); i++ {
		if tb1[i] > tb2[i] {
			return 1
		}
		if tb1[i] < tb2[i] {
			return -1
		}
	}

	return 0
}

func BestHand(hands []string) ([]string, error) {
	if len(hands) == 0 {
		return []string{}, nil
	}

	parsedHands := make([]Hand, len(hands))
	for i, handStr := range hands {
		hand, err := parseHand(handStr)
		if err != nil {
			return nil, fmt.Errorf("invalid hand %q: %v", handStr, err)
		}
		parsedHands[i] = hand
	}

	// Find best hand(s)
	best := []Hand{parsedHands[0]}
	for i := 1; i < len(parsedHands); i++ {
		cmp := compareHands(&parsedHands[i], &best[0])
		if cmp > 0 {
			best = []Hand{parsedHands[i]}
		} else if cmp == 0 {
			best = append(best, parsedHands[i])
		}
	}

	result := make([]string, len(best))
	for i, hand := range best {
		result[i] = hand.original
	}
	return result, nil
}
