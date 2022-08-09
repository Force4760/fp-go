package fp

// Take a slice of n elements from the original slice.
// If there are less than n elements, the original slice will be returned
func Take[T any](n int) func([]T) []T {
	return func(xs []T) []T {
		if n >= len(xs) {
			return xs
		}

		return xs[:n]
	}
}

// Take the elements of an array while it meet the condition specified in a callback function.
func TakeWhile[T any](predicate func(T) bool) func([]T) []T {
	return func(xs []T) []T {
		n := 0

		for _, x := range xs {
			if !predicate(x) {
				return xs[:n]
			}
			n++
		}

		return xs
	}
}

// See TakeWhile but callback receives index of element.
func TakeWhileWithIndex[T any](predicate func(T, int) bool) func([]T) []T {
	return func(xs []T) []T {
		n := 0

		for i, x := range xs {
			if !predicate(x, i) {
				return xs[:n]
			}
			n++
		}

		return xs
	}
}

// See Filter but callback receives index of element and the whole array.
func TakeWhileWithSlice[T any](predicate func(T, int, []T) bool) func([]T) []T {
	return func(xs []T) []T {
		n := 0

		for i, x := range xs {
			if !predicate(x, i, xs) {
				return xs[:n]
			}
			n++
		}

		return xs
	}
}
