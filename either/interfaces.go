package either

import "fmt"

// Implementation of important interfaces

// Implement the Stringer Interface
func (e Either[L, R]) String() string {
	if IsLeft(e) {
		return fmt.Sprint("Left(", e.left, ")")
	}

	return fmt.Sprint("Right(", e.right, ")")
}

func String[L, R any](e Either[L, R]) string {
	return e.String()
}
