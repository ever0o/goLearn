package main

import (
	"fmt"
	"unsafe"
)

/*
	常量
	常量是一个简单值得标识符，在程序运行的时候，不会被修改的量
	常量中的数据类型只可以是布尔型，数字型（包含整数，浮点和复数），字符串
*/

//显式类型常量的定义格式
const demoConst string = "Alex"

//隐式常量类型的定义格式
const demoConst2 = "Bob"

//常量用作为枚举
const (
	demoA = "abc"
	demoB = len(demoA)
	demoC = unsafe.Sizeof(demoA)
)

func main() {
	//常量的应用
	const LIGHT int = 10
	const WIDTH int = 5
	var area int

	area = LIGHT * WIDTH

	fmt.Printf("面积为 : %d", area)
	fmt.Println()
	fmt.Println(demoConst)
	fmt.Println(demoConst2)

	fmt.Println(demoA, demoB, demoC)
}
