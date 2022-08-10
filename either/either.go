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

// mapR :: (R -> T) -> Either L R -> Either L T
// mapL :: (L -> T) -> Either L R -> Either T R
// flatMapR :: (R -> Either L T) -> Either L R -> Either L T
// flatMapL :: (L -> Either T R) -> Either L R -> Either T R
// flat
// getROrElse :: R -> Either L R -> R
// getLOrElse :: L -> Either L R -> L
// flip :: Either L R -> Either R L
// match :: (L -> T) -> (R -> T) -> Either L R -> T
// isLeftAnd :: (L -> Bool) -> Either L R -> Bool
// isRightAnd :: (R -> Bool) -> Either L R -> Bool
