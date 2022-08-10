package fp

import (
	opt "github.com/repeale/fp-go/option"
	"github.com/repeale/fp-go/pair"
)

// ( head )(---- tail -- )
// [ 1,    2,    3,    4 ]
// ( --- init --- )( last )

// Get the head of the slice (the first element).
// Panics if the list is empty.
func Head[T any](xs []T) T {
	return xs[0]
}

// Get the head of the slice (the first element).
// Returns None if the slice is empty.
func SafeHead[T any](xs []T) opt.Option[T] {
	if IsNull(xs) {
		return opt.None[T]()
	}

	return opt.Some(xs[0])
}

// Get the tail of the slice (everything but the first element)
//
// If the argument is empty the result will also be empty ( [] -> [] )
func Tail[T any](xs []T) []T {
	if len(xs) == 0 {
		return []T{}
	}
	return xs[1:]
}

// Get the last element of slice.
// Panics if the list is empty.
func Last[T any](xs []T) T {
	return xs[len(xs)-1]
}

// Get the last element of slice as an Option.
// Returns None if the slice is empty.
func SafeLast[T any](xs []T) opt.Option[T] {
	if IsNull(xs) {
		return opt.None[T]()
	}

	return opt.Some(xs[len(xs)-1])
}

// Get the init of the slice (everything but the last)
//
// If the argument is empty the result will also be empty ( [] -> [] )
func Init[T any](xs []T) []T {
	length := len(xs)

	if length == 0 {
		return []T{}
	}

	return xs[:length-1]
}

// Check if a slice has no elements.
func IsNull[T any](xs []T) bool {
	return len(xs) == 0
}

// Check if all elements of a boolean slice are true.
func And(xs []bool) bool {
	return Every(Id[bool])(xs)
}

// Check if any element of a boolean slice is true.
func Or(xs []bool) bool {
	return Some(Id[bool])(xs)
}

// Check if a value is an element of a slice
func Elem[T comparable](e T) func(xs []T) bool {
	return Some(func(a T) bool { return a == e })
}

// Get the sum of all values of a slice
func Sum[T Num](xs []T) T {
	return Reduce(func(a, b T) T { return a + b }, 0)(xs)
}

// Get the product of all values of a slice
func Prod[T Num](xs []T) T {
	return Reduce(func(a, b T) T { return a * b }, 1)(xs)
}

// Get a slice with n elements equal to val
func Replicate[T any](n int) func(T) []T {
	return func(val T) []T {
		result := make([]T, n)

		for i := 0; i < n; i++ {
			result[i] = val
		}

		return result
	}
}

// Get a pair where the first element is a prefix of length n and the second element is the remainder of the slice.
func SplitAt[T any](n int) func([]T) pair.Pair[[]T, []T] {
	return func(xs []T) pair.Pair[[]T, []T] {
		if n == 0 {
			return pair.New([]T{}, xs)
		}

		if n >= len(xs) {
			return pair.New(xs, []T{})
		}

		return pair.New(xs[:n], xs[n:])
	}
}

// Concatenate two slices
func Concat[T any](xs []T) func([]T) []T {
	return func(ys []T) []T {
		return append(xs, ys...)
	}
}
