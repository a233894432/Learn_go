/*
Project:myexample
Time:2016年10月11日17:05:00
Author:diogo
*/

/*

2.4.3 Range

类似迭代器操作，返回 (索引, 值) 或 (键, 值)。

*/

package main

func main() {
	s := []int{1, 2, 3, 4, 5}

	for i, v := range s { // 复制 struct slice { pointer, len, cap }。

		if i == 0 {
			s = s[:3]  // 对 slice 的修改，不会影响 range。
			s[2] = 100 // 对底层数据的修改。
		}

		println(i, v)
		/*out
		  0 1
		  1 2
		  2 100
		  3 4
		  4 5
		*/
	}

}
