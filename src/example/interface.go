package main

import "fmt"

type USB interface {
	Name() string
	Connecter
}
type Connecter interface {
	Connect()
}

type PhoneConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() string  {
	return pc.name
}
func (pc PhoneConnecter) Connect()  {
	fmt.Println("Connect::",pc.name)
}

func  main()  {
	var a USB
	a=PhoneConnecter{"PhoneConnecter"}

	a.Connect()
	Disconnect(a)
}

func Disconnect(usb interface{})  {

	//if pc,ok:=usb.(PhoneConnecter);ok{
	//	fmt.Println("Disconnect",pc.name);
	//	return
	//}
	//
	//fmt.Println("Unknown Decive")

	switch v:=usb.(type) {
	case PhoneConnecter:
		fmt.Println("Disconnect",v.name);
	default:
		fmt.Println("Unknown Decive")


	}
}