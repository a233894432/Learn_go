/*
Project:Learn_go
Time:2016/8/20-17:32
author:Dio
*/
package main
import (
	"fmt"
	"crypto/aes"
	"strings"
)
func main(){
	rb:=[]byte {1,2,3,4,5,6,7,8,9,0,1,2,3,4,5,6};
	b:=make([]byte,16);
	strings.NewReader("1234567890123456").Read(b);
	// b=b[0:16];
	fmt.Printf("b: %s \n",b);
	cip,_:= aes.NewCipher(b);
	fmt.Println("cip:",cip,"err:");
	out:=make([]byte,len(rb));
	cip.Encrypt (rb, out);

}