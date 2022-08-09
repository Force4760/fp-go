package fp

// ( head )(---- tail -- )
// [ 1,    2,    3,    4 ]
// ( --- init --- )( last )

// Get the head of the slice (the first element).
// Panics if the list is empty.
func Head[T any](xs []T) T {
	return xs[0]
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
