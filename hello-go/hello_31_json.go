// run: go run hello_31_json.go
//
// encoding/json converts between Go values and JSON text. Struct tags
// (`json:"name"`) control the field name used on the wire; only exported
// fields are visible to json.Marshal/Unmarshal at all.
//
// Step 31: encoding/json
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"` // Step 31: tag controls the JSON key
	Age  int    `json:"age"`
}

func main() {
	p := Person{Name: "Ada", Age: 30}

	data, err := json.Marshal(p) // Step 31: Go value -> JSON bytes
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	var decoded Person
	if err := json.Unmarshal(data, &decoded); err != nil { // Step 31: JSON -> Go value
		panic(err)
	}
	fmt.Println(decoded)
}
