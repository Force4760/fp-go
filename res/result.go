package res

// Base Result struct.
type Result[T any] struct {
	isOk bool
	ok   T
	err  string
}

// Construcor for Result with Ok value.
func Ok[T any](ok T) Result[T] {
	return Result[T]{isOk: true, ok: ok}
}

// Construcor for Result with Err value.
func Err[T any](err string) Result[T] {
	return Result[T]{isOk: false, err: err}
}

// Check if a given Result is Ok.
func IsOk[T any](r Result[T]) bool {
	return r.isOk
}

// Check if a given Result is Err.
func IsErr[T any](r Result[T]) bool {
	return !r.isOk
}

// Check if a given Result is Ok.
// Then test the ok value to check if it satisfies a predicate.
func IsOkAnd[T any](pred func(T) bool) func(Result[T]) bool {
	return func(r Result[T]) bool {
		return r.isOk && pred(r.ok)
	}
}

// Execute the function on the Ok value, if it exists.
// Otherwise return the Err itself
func Map[T, R any](fn func(T) R) func(Result[T]) Result[R] {
	return func(r Result[T]) Result[R] {
		if IsErr(r) {
			return Err[R](r.err)
		}

		return Ok(fn(r.ok))
	}
}

// Removes one level of nesting of the Result.
func Flat[T any](r Result[Result[T]]) Result[T] {
	if IsErr(r) {
		return Err[T](r.err)
	}

	return r.ok
}

// Extracts the Ok value out of the Result if it exists.
// Otherwise returns a default value.
func GetOrElse[T any](def T) func(Result[T]) T {
	return func(r Result[T]) T {
		if IsErr(r) {
			return def
		}

		return r.ok
	}
}

// Check two Result for equality.
// The types must be comparable
func Eq[T comparable](r1 Result[T]) func(Result[T]) bool {
	return func(r2 Result[T]) bool {
		if IsOk(r1) && IsOk(r2) {
			return r1.ok == r2.ok
		}

		return false
	}
}

// Execute a function that returns a Result on the Ok value if it exists.
// Otherwise return the Err itself.
func Chain[T, R any](fn func(T) Result[R]) func(Result[T]) Result[R] {
	return func(r Result[T]) Result[R] {
		if IsErr(r) {
			return Err[R](r.err)
		}

		return fn(r.ok)
	}
}

// Extracts the ok value out of the Result with a function if it exists.
// Otherwise, extracts the Err value with another function.
func Match[T, R any](onErr func(string) R, onOk func(T) R) func(Result[T]) R {
	return func(r Result[T]) R {
		if IsErr(r) {
			return onErr(r.err)
		}

		return onOk(r.ok)
	}
}
