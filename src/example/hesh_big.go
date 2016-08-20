/*
Project:Learn_go
Time:2016/8/20-17:57
author:Dio
*/
package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/big"
)

func main() {
	bi := big.NewInt(0)
	h := md5.New()
	h.Write([]byte("diogoxiang"))
	hexstr := hex.EncodeToString(h.Sum(nil))
	bi.SetString(hexstr, 16)
	fmt.Println(bi.String())
}
