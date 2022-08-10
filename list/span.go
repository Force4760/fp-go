package lst

import "github.com/repeale/fp-go/pair"

// Get a Pair of arrays, where the first is the longest prefix of the provided array
// that satisfies a callback and the second is the rest of the array.
//
// span p <=> (takeWhile p, dropWhile p)
func Span[T any](pred func(T) bool) func([]T) pair.Pair[[]T, []T] {
	return func(xs []T) pair.Pair[[]T, []T] {
		n := 0

		for _, x := range xs {
			if !pred(x) {
				return pair.New(xs[:n], xs[n:])
			}
			n++
		}

		return pair.New(xs, []T{})
	}
}

// See Span but callback receives index of element.
func SpanWithIndex[T any](pred func(T, int) bool) func([]T) pair.Pair[[]T, []T] {
	return func(xs []T) pair.Pair[[]T, []T] {
		n := 0

		for i, x := range xs {
			if !pred(x, i) {
				return pair.New(xs[:n], xs[n:])
			}
			n++
		}

		return pair.New(xs, []T{})
	}
}

// See Span but callback receives index of element and the whole array.
func SpanWithSlice[T any](pred func(T, int, []T) bool) func([]T) pair.Pair[[]T, []T] {
	return func(xs []T) pair.Pair[[]T, []T] {
		n := 0

		for i, x := range xs {
			if !pred(x, i, xs) {
				return pair.New(xs[:n], xs[n:])
			}
			n++
		}

		return pair.New(xs, []T{})
	}
}
