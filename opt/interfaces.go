package opt

import "fmt"

// Implementation of important interfaces

// Implement the Stringer Interface
func (o Option[T]) String() string {
	if IsNone(o) {
		return "None()"
	}

	return fmt.Sprint("Some(", o.value, ")")
}

func String[T any](o Option[T]) string {
	return o.String()
}
