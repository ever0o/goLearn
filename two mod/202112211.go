package main

import "fmt"

/*
	自定义函数

	注：默认情况下，go语言使用的是值传递，在调用的过程中不会影响到实际参数
*/
func main() {

	var demoA int = 100
	var demoB int = 200

	var ret int

	ret = max(demoA, demoB)

	fmt.Printf("ret is : %d \n", ret)

	a, b := swap("Alex", "Bob")
	fmt.Printf("a is %s \nb is %s \n", a, b)
}

func max(num1, num2 int) int {
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

/*
	多个返回值的函数
*/
func swap(x, y string) (string, string) {
	return x, y
}
