package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	ch <- 1
	close(ch)
	receive1(ch)
}

func receive1(ch chan int) {
	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("通道已关闭")
			break
		}
		fmt.Printf("v:%#v ok:%#v", v, ok)
	}
}
