package main

import "reflect"

func main() {
	address := "www.w3cSchool.cn"

	//反射修改值必须要通过传递变量地址来修改。若函数传递的参数是值拷贝，则会发生以下错误：
	//panic: reflect: reflect.Value.SetString using unaddressable value
	//reflectSetValue1(address)
	reflectSetValue2(&address)
	println(address)
}

func reflectSetValue1(x interface{}) {
	value := reflect.ValueOf(x)
	if value.Kind() == reflect.String {
		value.SetString("欢迎来到W3Cschool")
	}
}

func reflectSetValue2(x interface{}) {
	value := reflect.ValueOf(x)
	if value.Elem().Kind() == reflect.String {
		value.Elem().SetString("欢迎来到W3CSchool")
	}
}
