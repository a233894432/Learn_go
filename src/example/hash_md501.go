/*
Project:Learn_go
Time:2016/8/20-16:42
author:Dio
*/
package main
import (
	"crypto/md5"
	"fmt"
	"encoding/hex"
)

func main(){
	md5Ctx := md5.New()
	md5Ctx.Write([]byte("test md5 encrypto"))
	cipherStr := md5Ctx.Sum(nil)
	fmt.Print(cipherStr)
	fmt.Print("\n")
	fmt.Print(hex.EncodeToString(cipherStr))
}
