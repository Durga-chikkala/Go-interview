// You can edit this code!
// Click here and start typing.
package main

import "fmt"

// List of unique items where the order is not mandatory.
type Set[T comparable] struct {
	set map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{set: make(map[T]struct{})}
}

func (s *Set[T]) Add(element T) {
	s.set[element] = struct{}{}
}

func (s *Set[T]) Delete(element T) {
	delete(s.set, element)
}

func (s *Set[T]) ViewSet() {
	var set []T
	for i, _ := range s.set {
		set = append(set, i)
	}

	fmt.Println(set)
}

func main() {
	s := NewSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(6)
	s.Add(2)
	s.Add(3)

	s.ViewSet()

}
