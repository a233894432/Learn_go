/*
Project:ch16
Time:2016/9/7-17:59
Author:diogo
*/
package main
import ( "fmt"

)
type nexter interface {
	next() byte
}
func nextFew1(n nexter, num int) []byte {
	var b []byte
	for i:=0; i < num; i++ {
		b[i] = n.next()
	}
	return b
}


func main() {


	fmt.Println("hello word")
}