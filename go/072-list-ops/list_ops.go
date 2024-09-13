package listops

type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	for _, num := range s {
		initial = fn(initial, num)
	}

	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	for i := len(s) - 1; i >= 0; i-- {
		initial = fn(s[i], initial)
	}
	return initial
}

func (s IntList) Filter(fn func(int) bool) IntList {
	result := make(IntList, 0, len(s))

	for _, num := range s {
		if fn(num) {
			result = append(result, num)
		}
	}

	return result
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	result := make(IntList, 0, len(s))

	for _, num := range s {
		result = append(result, fn(num))
	}

	return result
}

func (s IntList) Reverse() IntList {
	result := make(IntList, 0, len(s))

	for i := range s {
		result = append(result, s[len(s)-i-1])
	}

	return result
}

func (s IntList) Append(lst IntList) IntList {
	return append(s, lst...)
}

func (s IntList) Concat(lists []IntList) IntList {
	result := s
	for _, list := range lists {
		result = append(result, list...)
	}

	return result
}
