package main

import "fmt"

/*
	go 语言函数方法
	go 中区分函数和方法
*/
func main() {
	var c1 Circle
	c1.radius = 10
	fmt.Println("Area of Circle(c1) = ", c1.getArea())
}

type Circle struct {
	radius float64
}

//该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}
