/*
    author: diogoxiang
    time: 2016-08-16/8/22 - 22:16
    name: Learn_go
 */
package main

import (
	"encoding/gob"
	"log"
	"os"
	"fmt"
)

type Address struct {
	Type             string
	City             string
	Country          string
}

type VCard struct {
	FirstName   string
	LastName    string
	Addresses   []*Address
	Remark      string
}

var content string

func main() {
	pa := &Address{"private", "Aartselaar","Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa,wa}, "none"}
	fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
	// using a
	// n encoder:
	file, _ := os.OpenFile("src/example/vcard.gob", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := gob.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding gob")
	}
}
