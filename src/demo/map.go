package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 11
	fmt.Println("map:", m)
	//==	map: map[k1:7 k2:11]
	v := m["k1"]
	fmt.Println("v==", v)
	//== v==7
	delete(m, "k2")
	fmt.Println("map:", m)
	/*
		The optional second return value when getting a value from a map
		indicates if the key was present in the map. This can be used to
		disambiguate between missing keys and keys with zero values like 0 or "".
		Here we didn’t need the value itself, so we ignored it with the
		blank identifier _.
	*/
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	/*You can also declare and initialize a new map in the same line with this syntax.*/
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}
