package main

import "fmt"

/*
	go 语言类型转换
	注：go语言不支持隐式的格式转换
*/
func main() {

	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum) / float32(count)
	fmt.Printf("meain is %f\n", mean)
}
