package pov

type Tree struct {
	value    string
	children []*Tree
}

// New creates and returns a new Tree with the given root value and children.
func New(value string, children ...*Tree) *Tree {
	return &Tree{
		value:    value,
		children: children,
	}
}

// Value returns the value at the root of a tree.
func (tr *Tree) Value() string {
	return tr.value
}

// Children returns a slice containing the children of a tree.
// There is no need to sort the elements in the result slice,
// they can be in any order.
func (tr *Tree) Children() []*Tree {
	return tr.children
}

// String describes a tree in a compact S-expression format.
// This helps to make test outputs more readable.
// Feel free to adapt this method as you see fit.
func (tr *Tree) String() string {
	if tr == nil {
		return "nil"
	}
	result := tr.Value()
	if len(tr.Children()) == 0 {
		return result
	}
	for _, ch := range tr.Children() {
		result += " " + ch.String()
	}
	return "(" + result + ")"
}

// POV problem-specific functions

// findPath finds the path from the root to a node with the given value.
// Returns the path as a slice of trees, or nil if not found.
func (tr *Tree) findPath(target string) []*Tree {
	if tr.value == target {
		return []*Tree{tr}
	}
	for _, child := range tr.children {
		if path := child.findPath(target); path != nil {
			return append([]*Tree{tr}, path...)
		}
	}
	return nil
}

// reparent creates a new tree by reparenting based on the path.
// The path is from root to target, and we reverse the relationship.
func reparent(path []*Tree) *Tree {
	if len(path) == 0 {
		return nil
	}

	// Start with the target node (last in path)
	targetIdx := len(path) - 1
	target := path[targetIdx]

	// Create a new tree with target as root
	newRoot := &Tree{
		value:    target.value,
		children: make([]*Tree, 0),
	}

	// Add target's original children
	for _, child := range target.children {
		newRoot.children = append(newRoot.children, copyTree(child))
	}

	// Walk up the path, adding each parent as a child
	current := newRoot
	for i := targetIdx - 1; i >= 0; i-- {
		parent := path[i]
		newParent := &Tree{
			value:    parent.value,
			children: make([]*Tree, 0),
		}

		// Add parent's children (except the one in our path)
		nextInPath := path[i+1]
		for _, child := range parent.children {
			if child.value != nextInPath.value {
				newParent.children = append(newParent.children, copyTree(child))
			}
		}

		current.children = append(current.children, newParent)
		current = newParent
	}

	return newRoot
}

// copyTree creates a deep copy of a tree.
func copyTree(tr *Tree) *Tree {
	if tr == nil {
		return nil
	}
	newTree := &Tree{
		value:    tr.value,
		children: make([]*Tree, len(tr.children)),
	}
	for i, child := range tr.children {
		newTree.children[i] = copyTree(child)
	}
	return newTree
}

// FromPov returns the pov from the node specified in the argument.
func (tr *Tree) FromPov(from string) *Tree {
	path := tr.findPath(from)
	if path == nil {
		return nil
	}
	return reparent(path)
}

// PathTo returns the shortest path between two nodes in the tree.
func (tr *Tree) PathTo(from, to string) []string {
	// First, reparent from "from" node
	newTree := tr.FromPov(from)
	if newTree == nil {
		return nil
	}

	// Find path from new root to "to" node
	path := newTree.findPath(to)
	if path == nil {
		return nil
	}

	// Extract values from path
	result := make([]string, len(path))
	for i, node := range path {
		result[i] = node.value
	}
	return result
}
