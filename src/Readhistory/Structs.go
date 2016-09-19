// Go's _structs_ are typed collections of fields.
// They're useful for grouping data together to form
// records.

package main

import "fmt"

// This `person` struct type has `name` and `age` fields.
type person struct {
	name string
	age  int
}

func main() {

	// This syntax creates a new struct.
	fmt.Println(person{"Bob", 20}) // {Bob 20}

	// You can name the fields when initializing a struct.
	fmt.Println(person{name: "Alice", age: 30}) // {Alice 30}

	// Omitted fields will be zero-valued.
	fmt.Println(person{name: "Fred"}) // {Fred 0}

	// An `&` prefix yields a pointer to the struct.
	fmt.Println(&person{name: "Ann", age: 40}) // &{Ann 40}

	// Access struct fields with a dot.
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name) // Sean

	// You can also use dots with struct pointers - the
	// pointers are automatically dereferenced.
	sp := &s
	fmt.Println(sp.age) // 50

	// Structs are mutable.
	sp.age = 51
	fmt.Println(sp.age) // 51
}

/* out
{Bob 20}
{Alice 30}
{Fred 0}
&{Ann 40}
Sean
50
51
*/
