package main

import "fmt"

/*
	go向函数传递数组
	向函数传递数组的时候，可以设置数组的大小，也可以不用设置数组的大小
*/
func main() {

	var balance = []int{1000, 2, 3, 17, 50}
	var avg float32

	avg = getAverage(balance, 5)
	fmt.Printf("avg is : %f", avg)
}

func getAverage(arr []int, size int) float32 {
	var i, sum int
	var avg float32

	for i = 0; i < size; i++ {
		sum += arr[i]
	}

	avg = float32(sum / size)

	return avg
}
