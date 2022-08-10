package pair

import "fmt"

// Implementation of important interfaces

// Implement the Stringer Interface
func (p Pair[A, B]) String() string {
	return fmt.Sprint("(", p.a, ", ", p.b, ")")
}

func String[A, B any](p Pair[A, B]) string {
	return p.String()
}
