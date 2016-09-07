/*
Project:base64_encode
Time:2016/8/30-11:27
Author:diogo
*/
package main
import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	data := "MWFBU0Qz"
	sDec,_ := b64.StdEncoding.DecodeString(string(data))
	fmt.Printf("%s \n",sDec)
	//fmt.Println(data)
}