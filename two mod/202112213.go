package main

import (
	"fmt"
	"math"
)

/*
	go 函数作为值测试
*/
func main() {

	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	fmt.Println(getSquareRoot(9))
}
