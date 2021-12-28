package main

import "fmt"

/*
	指向指针的指针
*/
func main() {

	var a int
	var ptr *int
	var pptr **int

	a = 3000

	ptr = &a
	pptr = &ptr

	fmt.Printf("a is : %d\n", a)
	fmt.Printf("ptr is : %d\n", *ptr)
	fmt.Printf("pptr is %d\n", **pptr)
}
