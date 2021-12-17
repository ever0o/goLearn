package main

import "fmt"

/*
	iota
	iota，特殊常量，可以认为是一个可以被编译器修改的常量。
	在每一个const关键字出现时，被重置为0，然后再下一个const出现之前，每出现一次iota，其所代表的数字会自动增加1。


*/

func main() {
	//常规写法
	const (
		demoA = iota
		demoB = iota
		demoC = iota
	)
	fmt.Println(demoA, demoB, demoC)

	//简略写法
	const (
		demoD = iota
		demoE
		demoF
	)
	fmt.Println(demoD, demoE, demoF)

	//iota 用法
	const (
		demoAA = iota
		demoBB
		demoCC
		demoDD = "HA"
		demoEE
		demoFF = 100
		demoGG
		demoHH = iota
		demoII
	)
	fmt.Println(demoAA, demoBB, demoCC, demoDD, demoEE, demoFF, demoGG, demoHH, demoII)

	//iota 用法2
	const (
		numA = 1 << iota
		numB = 3 << iota
		numC
		numD
	)
	fmt.Println(numA, numB, numC, numD)
}
