/*
Project:ch17_18
Time:2016/9/7-18:28
Author:diogo
*/
package main

import (
	"fmt"
	"bytes"
)

func main() {
	a:=[]string{"aaa","dbb"}
	//fmt.Println(a[0])

	var buffer bytes.Buffer

	buffer.WriteString(a[0])
	buffer.WriteString(a[1])

	fmt.Print(buffer.String(), "\n")
}


/**
通过 buffer 串联字符串

类似于 Java 的 StringBuilder 类。

在下面的代码段中，我们创建一个 buffer，

通过 buffer.WriteString(s) 方法将字符串 s 追加到后面，最后再通过 buffer.String() 方法转换为 string：

var buffer bytes.Buffer
for {
    if s, ok := getNextString(); ok { //method getNextString() not shown here
        buffer.WriteString(s)
    } else {
        break
    }
}
fmt.Print(buffer.String(), "\n")
 */