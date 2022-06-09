package main

import (
	"fmt"
	"reflect"
)

/**
go 语言的反射
*/
func main() {
	var booknum float32 = 6
	var isbook bool = true
	bookauthor := "www.w3cschool.cn"
	bookdetail := make(map[string]string)
	bookdetail["Go语言教程"] = "www.w3cschool.cn"
	fmt.Println(reflect.TypeOf(booknum))
	fmt.Println(reflect.TypeOf(isbook))
	fmt.Println(reflect.TypeOf(bookauthor))
	fmt.Println(reflect.TypeOf(bookdetail))
}
