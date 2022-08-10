package fp

// Calls the specified callback function for all the elements in an array. The return value of the callback function is the accumulated result, and is provided as an argument in the next call to the callback function.
func Reduce[T any, R any](fn func(R, T) R, acc R) func([]T) R {
	return func(xs []T) R {

		for _, x := range xs {
			acc = fn(acc, x)
		}

		return acc
	}
}

// See Reduce but callback receives index of element.
func ReduceWithIndex[T any, R any](fn func(R, T, int) R, acc R) func([]T) R {
	return func(xs []T) R {

		for i, x := range xs {
			acc = fn(acc, x, i)
		}

		return acc
	}
}

// Like Reduce but callback receives index of element and the whole array.
func ReduceWithSlice[T any, R any](fn func(R, T, int, []T) R, acc R) func([]T) R {
	return func(xs []T) R {

		for i, x := range xs {
			acc = fn(acc, x, i, xs)
		}

		return acc
	}
}
