package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet, but you can learn about
// them here: https://go.dev/tour/generics/1

type Slicer interface {
	int | string | []int
}

func Keep[T Slicer](list []T, filterFunc func(T) bool) []T {
	var result []T
	for _, v := range list {
		if filterFunc(v) {
			result = append(result, v)
		}
	}

	return result
}

func Discard[T Slicer](list []T, filterFunc func(T) bool) []T {
	var result []T
	for _, v := range list {
		if !filterFunc(v) {
			result = append(result, v)
		}
	}

	return result
}
