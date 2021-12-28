package main

import "fmt"

/*
	指针
*/
func main() {
	var a int = 10
	fmt.Printf("a print is : %x\n", &a)

	/*
		指针测试用例
		在指针类型前面加*来获取指针所指向的内容
		在go语言中取地址符是用符号&，放到一个变量前使用就会返回相应变量的内存地址
	*/
	var demoA int = 20 /* 声明实际变量 */
	var ip *int        /* 声明指针变量 */

	ip = &demoA /* 指针变量的存储地址 */

	fmt.Printf("a address is : %x\n", &demoA)
	/* 指针变量的存储地址 */
	fmt.Printf("ip address is : %x\n", ip)
	/* 使用指针访问值 */
	fmt.Printf("*ip value is : %d\n", *ip)

	/*
		空指针
		当一个指针被定义后没有分配到任何变量时，它的值为nil
	*/
	var ptr *int
	fmt.Printf("ptr value is : %x\n", ptr)

}
