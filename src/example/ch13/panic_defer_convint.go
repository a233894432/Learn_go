/*
    author: diogoxiang
    time: 2016-08-16/8/30 - 01:04
    name: Learn_go
 */
/*
写一个 ConvertInt64ToInt 函数把 int64 值转换为 int 值，
如果发生错误（提示：参见 4.5.2.1 节）就 panic。
然后在函数 IntFromInt64 中调用这个函数并 recover，返回一个整数和一个错误。请测试这个函数！
 */
package main

import (
	"fmt"
	"math"
)

func main() {
	l := int64(15000)
	if i, err := IntFromInt64(l); err!= nil {
		fmt.Printf("The conversion of %d to an int32 resulted in an error: %s", l, err.Error())
	} else {
		fmt.Printf("%d converted to an int32 is %d", l, i)
	}
	fmt.Println()
	l = int64(math.MaxInt32 + 15000)
	if i, err := IntFromInt64(l); err!= nil {
		fmt.Printf("The conversion of %d to an int32 resulted in an error: %s", l, err.Error())
	} else {
		fmt.Printf("%d converted to an int32 is %d", l, i)
	}
}

func ConvertInt64ToInt(l int64) int {
	if math.MinInt32 <= l && l <= math.MaxInt32 {
		return int(l)
	}
	panic(fmt.Sprintf("%d is out of the int32 range", l))
}

func IntFromInt64(l int64) (i int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	i = ConvertInt64ToInt(l)
	return i, nil
}
/* Output:
15000 converted to an int32 is 15000
The conversion of 2147498647 to an int32 resulted in an error: 2147498647 is out of the int32 range
*/