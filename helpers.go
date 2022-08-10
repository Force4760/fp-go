package fp

// Number types
// + - * /
type Num interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64 | ~complex64 | ~complex128
}

// Orderable types
// < <= > >= == !=
type Ord interface {
	Num | ~string
}

// Function that returns the value it gets
func Id[T any](x T) T { return x }

// Ternary operator implementation
func If[T any](cond bool, onTrue, onFalse T) T {
	if cond {
		return onTrue
	}

	return onFalse
}
