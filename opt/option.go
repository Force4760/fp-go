package opt

// Base option struct
type Option[T any] struct {
	value    T
	hasValue bool
}

// Constructor for Option with a value
func Some[T any](value T) Option[T] {
	return Option[T]{value, true}
}

// Constructor for Option without a value
func None[T any]() Option[T] {
	return Option[T]{}
}

// Constructor for Option from a pointer. nil pointer == None, otherwise Some
func FromPtr[T any](ptr *T) Option[T] {
	if ptr == nil {
		return None[T]()
	}

	return Some(*ptr)
}

// Helper to check if the Option has a value
func IsSome[T any](o Option[T]) bool {
	return o.hasValue
}

// Helper to check if the Option has a value and if that value satisfies a predicate
func IsSomeAnd[T any](pred func(T) bool) func(Option[T]) bool {
	return func(o Option[T]) bool {
		return o.hasValue && pred(o.value)
	}
}

// Helper to check if the Option is missing the value
func IsNone[T any](o Option[T]) bool {
	return !o.hasValue
}

// Extracts the value out of the Option, if it exists. Otherwise returns a default value
func GetOrElse[T any](onNone T) func(Option[T]) T {
	return func(o Option[T]) T {

		if IsNone(o) {
			return onNone
		}

		return o.value
	}
}

// Extracts the value out of the Option, if it exists. Otherwise panics
func Get[T any](option Option[T]) T {
	if IsNone(option) {
		panic("Can't extract a value out of None")
	}

	return option.value
}

// Extracts the value out of the Option as a pointer, if it exists. Otherwise returns a nil pointer
func ToPtr[T any](o Option[T]) *T {
	if IsNone(o) {
		return nil
	}

	return &o.value
}

// Extracts the value out of the Option, if it exists, with a function. Otherwise returns the function with a default value
func Match[T, R any](onNone func() R, onSome func(T) R) func(Option[T]) R {
	return func(o Option[T]) R {

		if IsNone(o) {
			return onNone()
		}

		return onSome(o.value)
	}
}

// Execute the function on the Option value if it exists. Otherwise return the empty Option itself
func Map[T, R any](fn func(T) R) func(Option[T]) Option[R] {
	return func(o Option[T]) Option[R] {

		if IsNone(o) {
			return None[R]()
		}

		return Some(fn(o.value))
	}
}

// Execute a function that returns an Option on the Option value if it exists. Otherwise return the empty Option itself
func Chain[A any, B any](fn func(A) Option[B]) func(Option[A]) Option[B] {
	return func(o Option[A]) Option[B] {

		if IsNone(o) {
			return None[B]()
		}

		return fn(o.value)
	}
}

// Execute a predicate on the Option value if it exists.
// If the result is false or the Option is empty, return the empty Option.
// Otherwise, return the option itself
func Filter[T any](fn func(T) bool) func(Option[T]) Option[T] {
	return func(o Option[T]) Option[T] {
		if IsNone(o) {
			return None[T]()
		}

		if !fn(o.value) {
			return None[T]()
		}

		return o
	}
}

// Removes one level of nesting at a time. Option[Option[T]] -> Option[T]
func Flat[T any](o Option[Option[T]]) Option[T] {
	if IsNone(o) {
		return None[T]()
	}

	return o.value
}

// Check two Option for equality. The type must be comparable
func Eq[T comparable](o1 Option[T]) func(Option[T]) bool {
	return func(o2 Option[T]) bool {
		if IsNone(o1) {
			return IsNone(o2)
		}

		if IsNone(o2) {
			return false
		}

		return o1.value == o2.value
	}
}
