package main

import "fmt"

/*
	判断语句

*/
func main() {
	/*
		if语句
	*/
	var demoA int = 10

	if demoA < 20 {
		fmt.Println("demoA 小于20")
	}
	fmt.Println("demoA : ", demoA)

	/*
		switch 可以用来判断interface变量中存储的变量类型
	*/
	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf("x type : %T", i)
	case int:
		fmt.Printf("x type is int")
	case float64:
		fmt.Printf("x type is float64")
	case func(int) float64:
		fmt.Printf("x type is func(int)")
	case bool, string:
		fmt.Printf("x type is bool or string")
	default:
		fmt.Printf("don't know")
	}

	/*
		select 语句
	*/

	var demoC1, demoC2, demoC3 chan int
	var demoI1, demoI2 int
	select {
	case demoI1 = <-demoC1:
		fmt.Printf("received ", demoI1, " from demoC1\n")
	case demoC2 <- demoI2:
		fmt.Printf("sent ", demoI2, " from demoC3\n")
	case demoi3, ok := (<-demoC3):
		if ok {
			fmt.Printf("received ", demoi3, " from demoC3\n")
		} else {
			fmt.Printf("demoC3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
}
