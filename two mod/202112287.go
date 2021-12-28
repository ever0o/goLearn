package main

import "fmt"

/*
	go语言接口
*/
func main() {
	var phone phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}

type phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("i'm nokia, i can call you!")
}

type IPhone struct {
}

func (iphone IPhone) call() {
	fmt.Println("i'm iphone, i can call you!")
}
