package fp

// Returns the elements of an array that meet the condition specified in a callback function.
func Filter[T any](pred func(T) bool) func([]T) []T {
	return func(xs []T) []T {

		result := []T{}

		for _, x := range xs {
			if pred(x) {
				result = append(result, x)
			}
		}

		return result
	}
}

// See Filter but callback receives index of element.
func FilterWithIndex[T any](pred func(T, int) bool) func([]T) []T {
	return func(xs []T) []T {

		result := []T{}

		for i, x := range xs {
			if pred(x, i) {
				result = append(result, x)
			}
		}

		return result
	}
}

// Like Filter but callback receives index of element and the whole array.
func FilterWithSlice[T any](pred func(T, int, []T) bool) func([]T) []T {
	return func(xs []T) []T {

		result := []T{}

		for i, x := range xs {
			if pred(x, i, xs) {
				result = append(result, x)
			}
		}

		return result
	}
}
