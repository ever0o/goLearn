package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	ch <- 1
	close(ch)
	receive2(ch)
}

func receive2(ch chan int) {
	for i := range ch {
		fmt.Printf("v:%v", i)
	}
}
