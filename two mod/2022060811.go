package main

import (
	"fmt"
	"sync"
)

var (
	x1  int64
	wg2 sync.WaitGroup
)

// 对全局变量x执行5000次加1操作
func add() {
	for i := 0; i < 5000; i++ {
		x1 = x1 + 1
	}
	wg2.Done()
}

func main() {
	wg2.Add(2)
	go add()
	go add()
	wg2.Wait()
	fmt.Println(x)
}
