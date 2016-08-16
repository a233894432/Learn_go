/*
Project:Learn_go
Time:2016/8/16-10:50
author:Dio
*/
/*
7.1.4 将数组传递给函数

把第一个大数组传递给函数会消耗很多内存。有两种方法可以避免这种现象：

传递数组的指针
使用数组的切片
接下来的例子阐明了第一种方法：
 */
package main
import "fmt"

func main() {
	array := [3]float64{7.0, 8.5, 9.1}

	x := Sum(&array) // Note the explicit address-of operator
	// to pass a pointer to the array
	fmt.Printf("The sum of the array is: %f", x)
}

func Sum(a *[3]float64) (sum float64) {
	for _, v := range a { // derefencing *a to get back to the array is not necessary!
		sum += v
	}
	return
}