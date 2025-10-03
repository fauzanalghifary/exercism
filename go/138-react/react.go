package react

// reactor manages all cells and coordinates updates
type reactor struct {
	cells []cellInterface
}

// cellInterface is internal interface for tracking dependencies
type cellInterface interface {
	Cell
	update()
	addDependent(cellInterface)
}

// inputCell represents a cell with a settable value
type inputCell struct {
	value      int
	dependents []cellInterface
}

// computeCell represents a cell with a computed value
type computeCell struct {
	value      int
	deps       []cellInterface
	compute1   func(int) int
	compute2   func(int, int) int
	dependents []cellInterface
	callbacks  []*callback
}

// callback stores a callback function and its active status
type callback struct {
	id       int
	fn       func(int)
	canceled bool
}

// canceler allows removing a callback
type canceler struct {
	cell *computeCell
	id   int
}

func (c *canceler) Cancel() {
	for i := range c.cell.callbacks {
		if c.cell.callbacks[i].id == c.id {
			c.cell.callbacks[i].canceled = true
			return
		}
	}
}

func (c *inputCell) Value() int {
	return c.value
}

func (c *inputCell) SetValue(value int) {
	if c.value == value {
		return
	}
	c.value = value
	c.propagateUpdate()
}

func (c *inputCell) update() {
	// Input cells don't need to recalculate
}

func (c *inputCell) addDependent(dep cellInterface) {
	c.dependents = append(c.dependents, dep)
}

func (c *inputCell) propagateUpdate() {
	// Use topological sort to update cells in correct order
	visited := make(map[cellInterface]bool)
	var sorted []cellInterface

	// Topological sort (DFS post-order)
	var visit func(cell cellInterface)
	visit = func(cell cellInterface) {
		if visited[cell] {
			return
		}
		visited[cell] = true

		// Visit dependents first (children in dependency graph)
		switch c := cell.(type) {
		case *computeCell:
			for _, dep := range c.dependents {
				visit(dep)
			}
		case *inputCell:
			for _, dep := range c.dependents {
				visit(dep)
			}
		}
		sorted = append(sorted, cell)
	}

	// Start DFS from all direct dependents
	for _, dep := range c.dependents {
		visit(dep)
	}

	// Reverse to get correct update order (dependencies before dependents)
	for i, j := 0, len(sorted)-1; i < j; i, j = i+1, j-1 {
		sorted[i], sorted[j] = sorted[j], sorted[i]
	}

	// Update all cells and track value changes
	changed := make(map[cellInterface]bool)
	for _, cell := range sorted {
		oldValue := cell.Value()
		cell.update()
		if cell.Value() != oldValue {
			changed[cell] = true
		}
	}

	// Fire callbacks only for cells that actually changed
	for _, cell := range sorted {
		if cc, ok := cell.(*computeCell); ok && changed[cell] {
			cc.fireCallbacks()
		}
	}
}

func (c *computeCell) Value() int {
	return c.value
}

func (c *computeCell) SetValue(value int) {
	// Compute cells cannot have their value set directly
	panic("Cannot set value of compute cell")
}

func (c *computeCell) update() {
	if c.compute1 != nil {
		c.value = c.compute1(c.deps[0].Value())
	} else if c.compute2 != nil {
		c.value = c.compute2(c.deps[0].Value(), c.deps[1].Value())
	}
}

func (c *computeCell) addDependent(dep cellInterface) {
	c.dependents = append(c.dependents, dep)
}

func (c *computeCell) AddCallback(fn func(int)) Canceler {
	id := len(c.callbacks)
	cb := &callback{id: id, fn: fn, canceled: false}
	c.callbacks = append(c.callbacks, cb)
	return &canceler{cell: c, id: id}
}

func (c *computeCell) fireCallbacks() {
	for _, cb := range c.callbacks {
		if !cb.canceled {
			cb.fn(c.value)
		}
	}
}

func New() Reactor {
	return &reactor{cells: []cellInterface{}}
}

func (r *reactor) CreateInput(initial int) InputCell {
	c := &inputCell{value: initial, dependents: []cellInterface{}}
	r.cells = append(r.cells, c)
	return c
}

func (r *reactor) CreateCompute1(dep Cell, compute func(int) int) ComputeCell {
	depCell := dep.(cellInterface)
	c := &computeCell{
		deps:       []cellInterface{depCell},
		compute1:   compute,
		dependents: []cellInterface{},
		callbacks:  []*callback{},
	}
	c.value = compute(depCell.Value())
	depCell.addDependent(c)
	r.cells = append(r.cells, c)
	return c
}

func (r *reactor) CreateCompute2(dep1, dep2 Cell, compute func(int, int) int) ComputeCell {
	depCell1 := dep1.(cellInterface)
	depCell2 := dep2.(cellInterface)
	c := &computeCell{
		deps:       []cellInterface{depCell1, depCell2},
		compute2:   compute,
		dependents: []cellInterface{},
		callbacks:  []*callback{},
	}
	c.value = compute(depCell1.Value(), depCell2.Value())
	depCell1.addDependent(c)
	depCell2.addDependent(c)
	r.cells = append(r.cells, c)
	return c
}
