package main

import "fmt"

func main() {

	a := make(chan int)
	a <- 10
	fmt.Println("发送成功")
}
