package main

import (
	"fmt"
	"sync"
)

func main() {
	wg.Add(1) //计数器+1
	go hello1()
	fmt.Println("欢迎来到编程狮")
	wg.Wait() //阻塞代码的运行，知道计算器为0
}

var wg sync.WaitGroup

func hello1() {
	fmt.Println("hello")
	defer wg.Done() // 计数器-1
}
