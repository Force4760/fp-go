package lst

// Determines whether all the members of an array satisfy the specified test.
func Every[T any](pred func(T) bool) func([]T) bool {
	return func(xs []T) bool {

		for _, x := range xs {
			if !pred(x) {
				return false
			}
		}

		return true
	}
}

// See Every but callback receives index of element.
func EveryWithIndex[T any](pred func(T, int) bool) func([]T) bool {
	return func(xs []T) bool {

		for i, x := range xs {
			if !pred(x, i) {
				return false
			}
		}

		return true
	}
}

// Like Every but callback receives index of element and the whole array.
func EveryWithSlice[T any](pred func(T, int, []T) bool) func([]T) bool {
	return func(xs []T) bool {

		for i, x := range xs {
			if !pred(x, i, xs) {
				return false
			}
		}

		return true
	}
}
