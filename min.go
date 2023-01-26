package main

import "golang.org/x/exp/constraints"

// START1 OMIT
func Min[T any](x, y T) T {
	if x < y {
		return x
	}
	return y
}
// END1 OMIT

// START1.5 OMIT
func Min[T int|int64|float64](x, y T) T {
	if x < y {
		return x
	}
	return y
}
// END1.5 OMIT

// START2 OMIT
type minConstraint interface {
	uintptr|
		int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 | 
		float32 | float64 |
		string
}

func Min[T minConstraint](x, y T) T {
	if x < y {
		return x
	}
	return y
}
// END2 OMIT

// START3 OMIT
type minConstraint interface {
	uintptr| 
		int | int8 | int16 | int32 | int64 | 
		uint | uint8 | uint16 | uint32 | uint64 | 
		float32 | float64 | 
		string
}

func Min[T minConstraint](x, y T) T {
	if x < y {
		return x
	}
	return y
}

type MyInt int

func main() {
	x, y := MyInt(1), MyInt(2)
	Min[MyInt](x, y)
}
// END3 OMIT

// START4 OMIT
type minConstraint interface {
	~uintptr |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 |
		~string
}

func Min[T minConstraint](x, y T) T {
	if x < y {
		return x
	}
	return y
}
// END4 OMIT

// START5 OMIT
func Min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}
// END5 OMIT


// START6 OMIT
type Ordered interface {
	Integer | Float | ~string
}

type Integer interface {
	Signed | Unsigned
}

type Float interface {
	~float32 | ~float64
}
// END6 OMIT

// START7 OMIT
func Equal[T comparable](x, y T) bool {
	return x == y
}

func Foo[Key comparable, Value any](map[Key]Value) {
	// ...
}
// END7 OMIT
