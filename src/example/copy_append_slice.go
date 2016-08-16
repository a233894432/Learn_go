/*
Project:Learn_go
Time:2016/8/16-15:26
author:Dio
*/
package main
import "fmt"

func main() {
	sl_from := []int{1, 2, 3}
	sl_to := make([]int, 10)

	n := copy(sl_to, sl_from)
	fmt.Println(sl_to)
	fmt.Printf("Copied %d elements\n", n) // n == 3

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6,22)
	fmt.Println(sl3)

	sd:="22"
	fmt.Println(int(sd))


}
/**
7.6.8 切片和垃圾回收

切片的底层指向一个数组，该数组的实际体积可能要大于切片所定义的体积。
只有在没有任何切片指向的时候，底层的数组内层才会被释放，这种特性有时会导致程序占用多余的内存。

示例 函数 FindDigits 将一个文件加载到内存，然后搜索其中所有的数字并返回一个切片。

var digitRegexp = regexp.MustCompile("[0-9]+")

func FindDigits(filename string) []byte {
    b, _ := ioutil.ReadFile(filename)
    return digitRegexp.Find(b)
}
这段代码可以顺利运行，但返回的 []byte 指向的底层是整个文件的数据。
只要该返回的切片不被释放，垃圾回收器就不能释放整个文件所占用的内存。换句话说，一点点有用的数据却占用了整个文件的内存。

想要避免这个问题，可以通过拷贝我们需要的部分到一个新的切片中：

func FindDigits(filename string) []byte {
   b, _ := ioutil.ReadFile(filename)
   b = digitRegexp.Find(b)
   c := make([]byte, len(b))
   copy(c, b)
   return c
}

 */