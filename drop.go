package fp

// Get a slice with n elements removed from the original slice.
// If there are less than n elements, the empty slice will be returned
func Drop[T any](n int) func([]T) []T {
	return func(xs []T) []T {
		if n >= len(xs) {
			return []T{}
		}

		return xs[n:]
	}
}

// Remove elements of an array while they meet the condition specified in a callback function.
func DropWhile[T any](pred func(T) bool) func([]T) []T {
	return func(xs []T) []T {
		n := 0

		for _, x := range xs {
			if !pred(x) {
				return xs[n:]
			}
			n++
		}

		return []T{}
	}
}

// See DropWhile but callback receives index of element.
func DropWhileWithIndex[T any](pred func(T, int) bool) func([]T) []T {
	return func(xs []T) []T {
		n := 0

		for i, x := range xs {
			if !pred(x, i) {
				return xs[n:]
			}
			n++
		}

		return []T{}
	}
}

// See DropWhile but callback receives index of element and the whole array.
func DropWhileWithSlice[T any](pred func(T, int, []T) bool) func([]T) []T {
	return func(xs []T) []T {
		n := 0

		for i, x := range xs {
			if !pred(x, i, xs) {
				return xs[n:]
			}
			n++
		}

		return []T{}
	}
}
