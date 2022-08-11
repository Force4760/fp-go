package res

import "fmt"

// Implementation of important interfaces

// Implement the Stringer Interface
func (r Result[T]) String() string {
	if IsErr(r) {
		return fmt.Sprint("Err(", r.err, ")")
	}

	return fmt.Sprint("Ok(", r.ok, ")")
}

func String[T any](r Result[T]) string {
	return r.String()
}

// Implement the Error Interface
func (r Result[T]) Error() string {
	if IsOk(r) {
		return "No Error"
	}

	return r.err
}

func Error[T any](r Result[T]) string {
	return r.Error()
}
