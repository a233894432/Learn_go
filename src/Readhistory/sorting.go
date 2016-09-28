/*
Project:Readhistory
Time:2016/9/28-14:59
Author:diogo
*/
// Go's `sort` package implements sorting for builtins
// and user-defined types. We'll look at sorting for
// builtins first.

package main

import "fmt"
import "sort"

func main() {

	// Sort methods are specific to the builtin type;
	// here's an example for strings. Note that sorting is
	// in-place, so it changes the given slice and doesn't
	// return a new one.
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)   // out: Strings: [a b c]

	// An example of sorting `int`s.
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints) // out: Ints:    [2 4 7]

	// We can also use `sort` to check if a slice is
	// already in sorted order.
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)  // out: Sorted:  true
}

