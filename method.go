package main

import "golang.org/x/exp/constraints"

// START OMIT
type MyInt int

// Compilation error: "Method cannot have type parameters"
func (i *MyInt) Add[T constraints.Integer](other T) {
	*i += MyInt(other)
}
// END OMIT