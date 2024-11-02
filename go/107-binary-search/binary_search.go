package binarysearch

func SearchInts(list []int, key int) int {
	left, right := 0, len(list)-1

	for left <= right {
		middle := left + (right-left)/2
		middleElement := list[middle]

		if middleElement == key {
			return middle
		} else if middleElement > key {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	return -1
}
