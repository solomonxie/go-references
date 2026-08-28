// run: go test -v hello_29_testing_test.go
//
// A file ending in _test.go is excluded from normal builds and only
// compiled by "go test". A func Test<Name>(t *testing.T) is a test case;
// t.Errorf reports a failure without stopping the other cases in the loop
// (unlike t.Fatalf, which stops the current test immediately).
//
// Step 29: testing
package main

import "testing"

func Add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) { // Step 29: table-driven test
	cases := []struct {
		a, b, want int
	}{
		{1, 1, 2},
		{2, 3, 5},
		{-1, 1, 0},
	}

	for _, c := range cases {
		got := Add(c.a, c.b)
		if got != c.want {
			t.Errorf("Add(%d, %d) = %d, want %d", c.a, c.b, got, c.want)
		}
	}
}
