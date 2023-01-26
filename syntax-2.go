package main

//START1 OMIT
//START2 OMIT
//START3 OMIT
func Foo[S interface{~[]E}, E interface{}](s S) {
	// ...
}
//END1 OMIT

func Foo[S interface{~[]E}, E any](s S) {
	// ...
}
//END2 OMIT

func Foo[S ~[]E, E any](s S) {
	// ...
}
//END3 OMIT

//START4 OMIT
type SliceConstraint[T any] interface {
	~[]T
}

func Foo[S SliceConstraint[E], E any](s S) {
	// ...
}
//END4 OMIT
