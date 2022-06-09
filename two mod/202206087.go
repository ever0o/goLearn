package main

import "fmt"

func main() {
	a := make(chan int, 1)
	a <- 10
	fmt.Println("发送成功")
}
