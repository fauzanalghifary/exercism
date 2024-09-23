package kindergarten

import (
	"fmt"
	"sort"
	"strings"
)

type Garden struct {
	children map[string][]string
}

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

var plantsMap = map[string]string{
	"G": "grass",
	"C": "clover",
	"R": "radishes",
	"V": "violets",
}

func NewGarden(diagram string, children []string) (*Garden, error) {

	if diagram[0] != '\n' {
		return nil, fmt.Errorf("diagram must start with a newline")
	}

	g := &Garden{
		children: make(map[string][]string),
	}

	sortedChildren := make([]string, len(children))
	copy(sortedChildren, children)
	sort.Strings(sortedChildren)

	rows := strings.Split(diagram, "\n")
	for i, row := range rows {
		if i == 0 {
			continue
		}

		if len(row) != len(children)*2 {
			return nil, fmt.Errorf("row %d has a different number of plants than children", i)
		}

		for j, child := range sortedChildren {
			if len(g.children[child]) > 2 {
				return nil, fmt.Errorf("duplicate child %s", child)
			}

			if plantsMap[string(row[j*2])] == "" {
				return nil, fmt.Errorf("invalid plant %s", string(row[j*2]))
			}

			g.children[child] = append(g.children[child], plantsMap[string(row[j*2])])
			g.children[child] = append(g.children[child], plantsMap[string(row[j*2+1])])
		}
	}
	return g, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	plants, ok := g.children[child]
	return plants, ok
}
