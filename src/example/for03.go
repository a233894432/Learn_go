//2016/8/12 9:47
package main

import "fmt"

func main() {
	i := 0
	for { //since there are no checks, this is an infinite loop
		if i >= 3 {
			break
		}
		//break out of this for loop when this condition is met
		fmt.Println("Value of i is:", i)
		i++
	}
	fmt.Println("A statement just after for loop.")
	/* out
	Value of i is: 0
	Value of i is: 1
	Value of i is: 2
	A statement just after for loop.

	*/

	for i := 0; i < 7; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Odd:", i)
	}
	/* out:
	Odd: 1
	Odd: 3
	Odd: 5
	*/
}
