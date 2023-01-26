package main

import "fmt"

// START1 OMIT
// START2 OMIT
func PrintSliceInts(s []int64) {
	for _, elem := range s {
		fmt.Println(elem)
	}
}
//END1 OMIT

func PrintSliceFloats(s []float64) {
	for _, elem := range s {
		fmt.Println(elem)
	}
}
//END2 OMIT

// START3 OMIT
func PrintSliceTypeParameter[T interface{}](s []T) { // HL
	for _, elem := range s {
		fmt.Println(elem)
	}
}
//END3 OMIT

// START4 OMIT
func PrintSlice[T any](s []T) {
	for _, elem := range s {
		fmt.Println(elem)
	}
}
//END4 OMIT

// START5 OMIT
func main() {
	ints := []int64{1, 2, 3}
	floats := []float64{1.5, -2, 12}
	
	PrintSlice[int64](ints)
	PrintSlice[float64](floats)
}
//END5 OMIT
