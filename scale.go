package main

import "golang.org/x/exp/constraints"

// START1 OMIT
func ScaleV1[E constraints.Integer](s []E, scale E) []E {
	res := make([]E, len(s))
	for i, e := range s {
		res[i] = e * scale
	}
	return res
}
// END1 OMIT




// START2 OMIT
func ScaleV2[S ~[]E, E constraints.Integer](s S, scale E) S {
	res := make(S, len(s))
	for i, e := range s {
		res[i] = e * scale
	}
	return res
}
// END2 OMIT

// STARTVEC OMIT
type Vector []int
// ENDVEC OMIT

//START1EXAMPLE OMIT
v := Vector([]int{1, 2, 3})
res := ScaleV1(v, 2)
// type(res) == []int // HL
//END1EXAMPLE OMIT

//START2EXAMPLE OMIT
v := Vector([]int{1, 2, 3})
res := ScaleV2(v, 2)
// type(res) == Vector // HL
//END2EXAMPLE OMIT
