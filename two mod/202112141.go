package main

import "fmt"

/*
		语言变量
*/

/*
变量声明的三种方式
*/
var firstType string = "str"
var secondType = "str"

func main() {
	fmt.Println(firstType)

	fmt.Println(secondType)

	//该声明方式只能声明局部变量，不带声明格式的只能出现在函数体中出现
	thirdType := "str"
	fmt.Println(thirdType)


}
