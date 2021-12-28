package main

import "fmt"

/*
	语言切片
	go语言切片是对数组的抽象
	切片初始化   make([]T, length, capacity)    length 初始长度    capacity 容量
*/
func main() {
	var numbers = make([]int, 3, 5)
	var numbers2 []int
	numbers3 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)
	printSlice(numbers2)
	printSlice(numbers3)

	if numbers2 == nil {
		fmt.Printf("切片是空的")
	}
	//打印原始切片
	fmt.Println("numbers3 == ", numbers3)
	//打印切片从1到4
	fmt.Println("numbers3[1:4] == ", numbers3[1:4])
	//默认下限为0
	fmt.Println("numbers3[:3] == ", numbers3[:3])
	//默认上限为len(s)
	fmt.Println("numbers3[4:] == ", numbers3[4:])

	numbers4 := make([]int, 0, 5)
	printSlice(numbers4)

	numbers5 := numbers3[:2]
	printSlice(numbers5)

	numbers6 := numbers3[2:5]
	printSlice(numbers6)

	var numbers7 []int
	printSlice(numbers7)

	//追加空切片
	numbers7 = append(numbers7, 0)
	printSlice(numbers7)

	numbers7 = append(numbers7, 1)
	printSlice(numbers7)

	numbers7 = append(numbers7, 2, 3, 4)
	printSlice(numbers7)

	numbers8 := make([]int, len(numbers7), (cap(numbers7))*2)

	copy(numbers8, numbers7)
	printSlice(numbers8)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
