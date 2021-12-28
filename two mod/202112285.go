package main

import "fmt"

/*
	go 语言递归函数
	使用递归来实现阶乘
*/
func main() {
	var i int = 15
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(i))

	var i2 int
	for i2 = 0; i2 < 10; i2++ {
		fmt.Printf("%d\t", Fibonacci(i2))
	}
}

func Factorial(x int) (result int) {
	if x == 0 {
		result = 1
	} else {
		result = x * Factorial(x-1)
	}
	return
}

func Fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return Fibonacci(n-2) + Fibonacci(n-1)
}
