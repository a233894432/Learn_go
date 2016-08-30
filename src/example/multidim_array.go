/*
Project:Learn_go
Time:2016/8/16-10:46
author:Dio
*/
/*
7.1.3 多维数组
 */
package main
const (
	WIDTH  = 1920
	HEIGHT = 1080
)

type pixel int
var screen [WIDTH][HEIGHT]pixel

func main() {
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			screen[x][y] = 0

		}
	}
}