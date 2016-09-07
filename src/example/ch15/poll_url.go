/*
Project:ch15
Time:2016/9/7-16:41
Author:diogo
*/
/**
在下边这个程序中，数组中的url都将被访问：会发送一个简单的http.Head()请求查看返回值；
它的声明如下：func Head(url string) (r *Response, err error)
*/
package main

import (
	"fmt"
	"net/http"
)

var urls = []string{
	"http://www.google.com/",
	"http://golang.org/",
	"http://blog.golang.org/",
}

func main() {
	// Execute an HTTP HEAD request for all url's
	// and returns the HTTP status string or an error string.
	for _, url := range urls {
		resp, err := http.Head(url)
		if err != nil {
			fmt.Println("Error:", url, err)
		}
		fmt.Print(url, ": ", resp.Status)
	}
}
