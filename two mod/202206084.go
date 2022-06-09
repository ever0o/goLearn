package main

import (
	"fmt"
	"sync"
)

var wg1 sync.WaitGroup

func main() {

	for i := 0; i < 10; i++ {
		go hello2(i)
		wg1.Add(1)
	}
	wg1.Wait()
}

func hello2(i int) {
	fmt.Printf("hello,欢迎来到编程狮%v\n", i)
	defer wg1.Done()
}
