package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(
		records, func(i, j int) bool {
			return records[i].ID < records[j].ID
		},
	)

	nodes := make(map[int]*Node)
	for i, record := range records {
		if record.ID == 0 && record.Parent != 0 {
			return nil, errors.New("invalid input: root node has parent")
		}

		if nodes[record.ID] != nil {
			return nil, errors.New("invalid input: duplicate node")
		}

		if i != record.ID {
			return nil, errors.New("invalid input: non-continuous")
		}

		if record.ID != 0 && record.ID == record.Parent {
			return nil, errors.New("invalid input: cycle directly")
		}

		if record.ID < record.Parent {
			return nil, errors.New("invalid input: cycle indirectly")
		}

		nodes[record.ID] = &Node{ID: record.ID}
		if record.ID != 0 {
			parent := nodes[record.Parent]
			parent.Children = append(parent.Children, nodes[record.ID])
		}
	}

	return nodes[0], nil
}
