package main

import "golang.org/x/exp/constraints"

// START OMIT
func Min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func main() {
	var m, x, y int
	m = Min[int](x, y)
	m = Min(x, y) // HL
	
	var mf, a, b float64 
	mf = Min(a, b) // HL
}
// END OMIT
