/*
Project:Learn_go
Time:2016/8/20-16:26
author:Dio
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "noneeE"}
	// fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
	// JSON format:
	js, _ := json.MarshalIndent(vc," ","  ")
	//格式化数据
	fmt.Printf("JSON format: %s \n", js)
	//单行数据
	js01,_:=json.Marshal(vc)
	fmt.Printf("JSON line: %s",js01)

	// using an encoder:
	file, _ := os.OpenFile("src/example/vcard.json", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := json.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding json")
	}
}
