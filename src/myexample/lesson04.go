/*
Project:myexample
Time:2016/9/30-14:26
Author:diogo
*/
package main

func main() {
//	配合标签，break 和 continue 可在多级嵌套循环中跳出。
L1:
	for x := 0; x < 3; x++ {
	L2:
		for y := 0; y < 5; y++ {
			if y > 2 {
				continue L2
			}
			if x > 1 {
				break L1
			}

			print(x, ":", y, " ")
		}

		println()
	}
}

/* OUT
0:0 0:1 0:2
1:0 1:1 1:2

 */