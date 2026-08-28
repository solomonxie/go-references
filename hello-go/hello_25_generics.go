// run: go run hello_25_generics.go
//
// A type parameter, written in [T Constraint], lets one function or type
// work over several concrete types while the compiler still checks it at
// compile time. A constraint is an interface listing which types are
// allowed.
//
// Step 25: generics
package main

import "fmt"

type Number interface { // Step 25: constraint = the set of allowed types
	int | float64
}

func Sum[T Number](nums []T) T { // Step 25: T is inferred from the call site
	var total T
	for _, n := range nums {
		total += n
	}
	return total
}

type Stack[T any] struct { // Step 25: generic struct, T fixed per instance
	items []T
}

func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last, true
}

func main() {
	fmt.Println(Sum([]int{1, 2, 3}))
	fmt.Println(Sum([]float64{1.5, 2.5}))

	var st Stack[string]
	st.Push("a")
	st.Push("b")
	v, _ := st.Pop()
	fmt.Println(v)
}
