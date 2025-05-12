/*
Package stream allows LINQ-like functions on slices. It is similar to lodash, C# LINQ, Java streams,
etc.. This can all be done through simple for loops, but this makes me feel all warm and fuzzy inside
and is a good exercise in understanding generics in Golang a bit better
*/
package stream

import (
	"errors"
)

// Stream is simply a container for a slice (S) of type T
type Stream[S ~[]T, T comparable] struct {
	contents []T
}

// NewStream creates a new Stream object with a slice (S) of type T
func NewStream[S ~[]T, T comparable](input S) *Stream[S, T] {
	var contents = make(S, len(input))
	for ii := range input {
		contents[ii] = input[ii]
	}
	return &Stream[S, T]{contents}
}

// Slice returns a copy of the stream's private slice
func (s *Stream[S, T]) Slice() S {
	cc := make([]T, len(s.contents))

	copy(cc, s.contents)
	return cc
}

// Filter prunes the inner slice to only those items fulfilling the given predicate
func (s *Stream[S, T]) Filter(predicate func(T) bool) *Stream[S, T] {
	s.contents = s.Find(predicate)

	return s
}

// Find returns all items matching the given predicate
func (s *Stream[S, T]) Find(predicate func(T) bool) S {
	res := make([]T, 0)
	for _, v := range s.contents {
		if predicate(v) {
			res = append(res, v)
		}
	}

	return res
}

// FindFirst returns the first item matching the given predicate. Error returned is nil if
// an item is found, or not nil if the item is not found
func (s *Stream[S, T]) FindFirst(filterFunction func(T) bool) (T, error) {
	for _, v := range s.contents {
		if filterFunction(v) {
			return v, nil
		}
	}

	var result T
	return result, errors.New("no element found")
}

// Length returns the length of the internal slice
func (s *Stream[S, T]) Length() int {
	return len(s.contents)
}

// Map performs the callback function to mutate every element in the internal slice
func (s *Stream[S, T]) Map(callback func(T) T) *Stream[S, T] {
	for ii := range s.contents {
		s.contents[ii] = callback(s.contents[ii])
	}

	return s
}

// Foreach performs the callback function using every element in the internal slice
func (s *Stream[S, T]) Foreach(callback func(int, T)) {
	for ii, v := range s.contents {
		callback(ii, v)
	}
}

// Reduce performs a reduce function on every element in the internal slice
func (s *Stream[S, T]) Reduce(reduceFunction func(T, T) T) T {
	var value T
	for _, v := range s.contents {
		value = reduceFunction(value, v)
	}

	return value
}
