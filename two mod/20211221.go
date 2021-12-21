package main

import "fmt"

/*
	循环控制语句
	break continue goto
*/
func main() {
	var demoA int = 10

	for demoA < 20 {
		fmt.Printf("secA => demoA is : %d \n", demoA)
		demoA++
		if demoA > 15 {
			//使用break语句跳出循环
			break
		}
	}
	fmt.Println()

	var demoB int = 10

	for demoB < 20 {
		if demoB == 15 {
			demoB = demoB + 1
			continue
		}
		fmt.Printf("secB => demoB is : %d \n", demoB)
		demoB++
	}

	fmt.Println()

	// goto语句通常与条件语句配合使用，可用来实现条件转移，构成循环，跳出循环体等功能
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if y == 2 {
				goto breakHere
			}
		}
	}
	return

breakHere:
	fmt.Println("done")
}
