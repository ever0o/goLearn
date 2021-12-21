package main

import "fmt"

/*
	自定义函数值传递的测试用例

	注  层级为 包下   各个go文件中的方法名不能重复
*/
func main() {
	var a int = 100
	var b int = 200

	fmt.Printf("a change before is %d\n", a)
	fmt.Printf("b change before is %d\n", b)
	swapB(a, b)
	fmt.Printf("a change after is %d\n", a)
	fmt.Printf("b change after is %d\n", b)

	swapC(&a, &b)
	fmt.Printf("&a change after is %d\n", a)
	fmt.Printf("&b change after is %d\n", b)
}

/*
	值传递测试
*/
func swapB(x, y int) int {
	var tmp int
	tmp = x
	y = x
	y = tmp

	return tmp
}

/*
	引用传递测试
*/
func swapC(x *int, y *int) {
	var tmp int
	tmp = *x
	*x = *y
	*y = tmp
}
