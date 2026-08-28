// run: go run hello_12_maps.go
//
// A map is a hash table: map[KeyType]ValueType. Reading a missing key
// returns the value type's zero value instead of erroring, so the
// "comma-ok" form is how code checks whether a key was actually present.
//
// Step 12: maps
package main

import "fmt"

func main() {
	ages := map[string]int{"Alice": 30, "Bob": 25} // Step 12: map literal
	ages["Carol"] = 40                             // Step 12: insert/update

	fmt.Println(ages)

	age, ok := ages["Dave"] // Step 12: comma-ok, Dave isn't in the map
	fmt.Println("Dave:", age, "present:", ok)

	delete(ages, "Bob") // Step 12: remove a key
	fmt.Println(ages)
}
