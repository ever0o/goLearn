package main

import "fmt"

/**
循环语句
*/
func main() {
	var demoA int
	var demoB int = 15

	numbers := [6]int{1, 2, 3, 5}

	for demoA := 0; demoA < 10; demoA++ {
		fmt.Printf("A demoA is : %d \n", demoA)
	}

	for demoA < demoB {
		demoA++
		fmt.Printf("B demoA is : %d \n", demoA)
	}

	/**
	range格式
	语句执行过程：
	for key, value := range oldMap {
		newMap[key] = value
	}

	*/
	for i, x := range numbers {
		fmt.Printf(" 第%d位 x 的值 = %d\n", i, x)
	}
}
