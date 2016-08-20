/*
Project:Learn_go
Time:2016/8/20-16:44
author:Dio
*/
package main
import (
	"fmt"
	"crypto/md5"
	"encoding/base64"
	"time"
)
var cipher = "DIO"
var h = md5.New()
func cipherEncode(sourceText string) string {
	h.Write([]byte(cipher))
	cipherHash := fmt.Sprintf("%x", h.Sum(nil))
	h.Reset()
	inputData := []byte(sourceText)
	loopCount := len(inputData)
	outData := make([]byte,loopCount)
	for i:= 0; i < loopCount ; i++ {
		outData[i] = inputData[i] ^ cipherHash[i%32]
	}
	return fmt.Sprintf("%s", outData)
}
func encode(sourceText string) string {
	h.Write([]byte(time.Now().Format("2006-01-02 15:04:05")))
	noise := fmt.Sprintf("%x", h.Sum(nil))
	h.Reset()
	inputData := []byte(sourceText)
	loopCount := len(inputData)
	outData := make([]byte,loopCount*2)

	for i, j := 0,0; i < loopCount ; i,j = i+1,j+1 {
		outData[j] = noise[i%32]
		j++
		outData[j] = inputData[i] ^ noise[i%32]
	}

	return base64.StdEncoding.EncodeToString([]byte(cipherEncode(fmt.Sprintf("%s", outData))))
}
func decode(sourceText string) string {
	buf, err := base64.StdEncoding.DecodeString(sourceText)
	if err != nil {
		fmt.Println("Decode(%q) failed: %v", sourceText, err)
		return ""
	}
	inputData := []byte(cipherEncode(fmt.Sprintf("%s", buf)))
	loopCount := len(inputData)
	outData := make([]byte,loopCount)
	for i, j := 0,0; i < loopCount ; i,j = i+2,j+1 {
		outData[j] = inputData[i] ^ inputData[i+1]
	}
	return fmt.Sprintf("%s", outData)
}
func main() {
	s := encode("張學友")
	d:=encode("向秀成")
	fmt.Println(s)

	fmt.Println(decode(s))
	fmt.Println(d)
	fmt.Println(decode(d))
}
