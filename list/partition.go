package lst

import "github.com/force4760/fp-go/pair"

// Get a Pair of arrays, where the first has all the elements that satisfy a callback
// and the second has all the elements that don't satisty it.
//
// partition p xs <=> (filter p xs, filter (not . p) xs)
func Partition[T any](pred func(T) bool) func([]T) pair.Pair[[]T, []T] {
	return func(xs []T) pair.Pair[[]T, []T] {
		res1 := []T{}
		res2 := []T{}

		for _, x := range xs {
			if pred(x) {
				res1 = append(res1, x)
			} else {
				res2 = append(res2, x)
			}
		}

		return pair.New(res1, res2)
	}
}

// See Partition but callback receives index of element.
func PartitionWithIndex[T any](pred func(T, int) bool) func([]T) pair.Pair[[]T, []T] {
	return func(xs []T) pair.Pair[[]T, []T] {
		res1 := []T{}
		res2 := []T{}

		for i, x := range xs {
			if pred(x, i) {
				res1 = append(res1, x)
			} else {
				res2 = append(res2, x)
			}
		}

		return pair.New(res1, res2)
	}
}

// See Partition but callback receives index of element and the whole array.
func PartitionWithSlice[T any](pred func(T, int, []T) bool) func([]T) pair.Pair[[]T, []T] {
	return func(xs []T) pair.Pair[[]T, []T] {
		res1 := []T{}
		res2 := []T{}

		for i, x := range xs {
			if pred(x, i, xs) {
				res1 = append(res1, x)
			} else {
				res2 = append(res2, x)
			}
		}

		return pair.New(res1, res2)
	}
}
