package main

import "sort"

// STARTSLICE OMIT
type Stack[T any] []T

func (s *Stack[T]) Push(val T) {
	*s = append(*s, val)
}

func (s *Stack[T]) Pop() T {
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val
}
// ENDSLICE OMIT

// STARTSORT1 OMIT
// STARTSORT2 OMIT
type SliceSort[T any] struct {
	s []T
	less func(T, T) bool
}

func (s SliceSort[T]) Len() int { return len(s.s) }
func (s SliceSort[T]) Swap(i, j int) { s.s[i], s.s[j] = s.s[j], s.s[i] }
func (s SliceSort[T]) Less(i, j int) bool { return s.less(s.s[i], s.s[j]) }
// ENDSORT1 OMIT

func Sort[T any](s []T, less func(T, T) bool) {
	sort.Sort(SliceSort[T]{s, less})
}
// ENDSORT2 OMIT
