/*
Project:myexample
Time:2016年10月11日14:15:16
Author:diogo
*/

package main

import "fmt"

/** out
字符串的类型转换
*/
func main() {
	s := "abc汉字"

	for i := 0; i < len(s); i++ { // byte
		fmt.Printf("%c,", s[i])
	}

	fmt.Println()

	for _, r := range s { // rune
		fmt.Printf("%c,", r)
	}
}

/* out:
a,b,c,,±,,,,,
a,b,c,汉,字,
*/
