package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}

	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	qty, exists := units[unit]
	if !exists {
		return false
	}

	if _, exists := bill[item]; exists {
		bill[item] += qty
	} else {
		bill[item] = qty
	}

	// another correct implementation
	//bill[item] += qty

	// WRONG implementation because it copies a value, not a reference.
	//billItem, exists := bill[item]
	//if exists {
	//	billItem += qty
	//} else {
	//	billItem = qty
	//}

	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	qty, exists := units[unit]
	if !exists {
		return false
	}

	existingQty, exists := bill[item]
	if !exists || qty > existingQty {
		return false
	}

	if qty == existingQty {
		delete(bill, item)
	} else {
		bill[item] -= qty
	}

	return true

}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	qty, exists := bill[item]
	return qty, exists
}
