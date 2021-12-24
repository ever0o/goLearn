package main

import "fmt"

/*
	go语言变量作用域
*/

/*
	全局变量
*/
var g int
var h int = 10

func main() {
	/*
		局部变量
	*/
	var a, b, c int
	var h int = 20
	var i int = 0

	a = 10
	b = 20
	c = a + b
	g = a + b + c

	fmt.Printf("result : a = %d, b = %d and c = %d \n", a, b, c)
	fmt.Printf("result : g = %d\n", g)
	// 在GO中，全局变量和局部变量的命名可以相同，但是函数内的局部变量会被优先考虑和使用
	fmt.Printf("result : h = %d\n", h)

	fmt.Printf("main() a = %d\n", a)
	i = sum(a, b)
	fmt.Printf("main() c = %d\n", i)
}

func sum(a, b int) int {
	fmt.Printf("sum() a = %d\n", a)
	fmt.Printf("sum() b = %d\n", b)
	return a + b
}
