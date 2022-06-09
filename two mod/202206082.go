package main

import (
	"fmt"
	"time"
)

/**
并发
*/
func main() {
	go hello()
	fmt.Println("欢迎来到编程狮")
	time.Sleep(time.Second)
}

func hello() {
	fmt.Println("hello")
}
