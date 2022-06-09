package main

import "fmt"

func main() {
	a := make(chan int)
	go receive(a)
	a <- 10
	fmt.Println("发送成功")
}

func receive(x chan int) {
	ret := <-x
	fmt.Println("接收成功", ret)
}
