package either

// Base Either struct.
type Either[L, R any] struct {
	isLeft bool
	left   L
	right  R
}

// Construcor for Either with Left value.
func Left[R, L any](val L) Either[L, R] {
	return Either[L, R]{isLeft: true, left: val}
}

// Construcor for Either with Right value.
func Right[L, R any](val R) Either[L, R] {
	return Either[L, R]{isLeft: false, right: val}
}

// Check if a given Either is Left.
func IsLeft[L, R any](e Either[L, R]) bool {
	return e.isLeft
}

// Check if a given Either is Right.
func IsRight[L, R any](e Either[L, R]) bool {
	return !e.isLeft
}

// Check if a given Either is Left.
// Then test the left value to check if it satisfies a predicate.
func IsLeftAnd[R, L any](fn func(L) bool) func(Either[L, R]) bool {
	return func(e Either[L, R]) bool {
		return e.isLeft && fn(e.left)
	}
}

// Check if a given Either is left.
// Then test the right value to check if it satisfies a predicate.
func IsRightAnd[L, R any](fn func(R) bool) func(Either[L, R]) bool {
	return func(e Either[L, R]) bool {
		return !e.isLeft && fn(e.right)
	}
}

// Extracts the left value out of the Either, if it exists. Otherwise returns a default value.
func GetOrElseL[R, L any](def L) func(Either[L, R]) L {
	return func(e Either[L, R]) L {
		if IsLeft(e) {
			return e.left
		}

		return def
	}
}

// Extracts the right value out of the Either, if it exists. Otherwise returns a default value.
func GetOrElseR[L, R any](def R) func(Either[L, R]) R {
	return func(e Either[L, R]) R {
		if IsRight(e) {
			return e.right
		}

		return def
	}
}

// Check two Either for equality. The types must be comparable
func Eq[L, R comparable](e1 Either[L, R]) func(Either[L, R]) bool {
	return func(e2 Either[L, R]) bool {
		if IsLeft(e1) {
			if IsLeft(e2) {
				return e1.left == e2.left
			}

			return false
		} else {
			if IsRight(e2) {
				return e1.right == e2.right
			}

			return false
		}
	}
}

// Switch the types of Either: Either[L, R] -> Either[R, L]
func Flip[L, R any](e Either[L, R]) Either[R, L] {
	if IsLeft(e) {
		return Right[R](e.left)
	}

	return Left[L](e.right)
}

// Extracts the value out of the Either with the onLeft function if it is Left and with the onRight function if it is Right
func Match[L, R, T any](onLeft func(L) T, onRight func(R) T) func(Either[L, R]) T {
	return func(e Either[L, R]) T {
		if IsLeft(e) {
			return onLeft(e.left)
		}

		return onRight(e.right)
	}
}

// Execute the function on the Left value if it exists. Otherwise return the Right itself
func MapL[R, L, T any](fn func(L) T) func(Either[L, R]) Either[T, R] {
	return func(e Either[L, R]) Either[T, R] {
		if IsLeft(e) {
			return Left[R](fn(e.left))
		}

		return Right[T](e.right)
	}
}

// Execute the function on the Right value if it exists. Otherwise return the Left itself
func MapR[L, R, T any](fn func(R) T) func(Either[L, R]) Either[L, T] {
	return func(e Either[L, R]) Either[L, T] {
		if IsRight(e) {
			return Right[L](fn(e.right))
		}

		return Left[T](e.left)
	}
}

// Extracts the left value of the Either with a function, if it exists. Otherwise returns the right value.
func ChainL[L, R, T any](fn func(L) Either[T, R]) func(Either[L, R]) Either[T, R] {
	return func(e Either[L, R]) Either[T, R] {
		if IsRight(e) {
			return Right[T](e.right)
		}

		return fn(e.left)
	}
}

// Extracts the right value of the Either with a function, if it exists. Otherwise returns the left value.
func ChainR[L, R, T any](fn func(R) Either[L, T]) func(Either[L, R]) Either[L, T] {
	return func(e Either[L, R]) Either[L, T] {
		if IsLeft(e) {
			return Left[T](e.left)
		}

		return fn(e.right)
	}
}

// Removes one level of nesting on the left value
func FlatL[L, R any](e Either[Either[L, R], R]) Either[L, R] {
	if IsRight(e) {
		return Right[L](e.right)
	}

	return e.left
}

// Removes one level of nesting on the right value
func FlatR[L, R any](e Either[L, Either[L, R]]) Either[L, R] {
	if IsLeft(e) {
		return Left[R](e.left)
	}

	return e.right
}
