package sublist

import (
	"slices"
)

func Sublist(l1, l2 []int) Relation {
	if len(l1) == len(l2) {
		return CompareList(l1, l2, RelationEqual)
	} else if len(l1) > len(l2) {
		return CompareList(l1, l2, RelationSuperlist)
	} else {
		return CompareList(l2, l1, RelationSublist)
	}
}

func CompareList(longArray, shortArray []int, related Relation) Relation {
	if len(shortArray) == 0 {
		return related
	}
	for i := range longArray {
		if longArray[i] == shortArray[0] {
			isValidIndex := i+len(shortArray) <= len(longArray)
			if isValidIndex && slices.Equal(shortArray, longArray[i:i+len(shortArray)]) {
				return related
			}
		}
	}
	return RelationUnequal
}
