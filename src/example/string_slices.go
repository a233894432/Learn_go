/*
Project:Learn_go
Time:2016/8/16-16:07
author:Dio
*/
/*
练习 7.12
*/
package main

import "fmt"

func main() {
	a := "tinglo"

	a1, a2 := TrimString(a, 3)
	fmt.Println("a1===", a1)
	fmt.Println("a2===", a2)
	/**
		练习 7.13

	 */
	stra:="diogoxiang"
	fmt.Println(stra[len(stra)/2:] + stra[:len(stra)/2]) // out : xiangdiogo


}

func TrimString(str string, i int) (string, string) {

	s := []byte(str)

	s1 := s[:i]
	s2 := s[i:]

	fmt.Printf("%s \n", s1)

	fmt.Printf("%s \n", s2)

	return string(s1), string(s2)

}
