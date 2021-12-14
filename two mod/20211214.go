package main

import (
	"fmt"
)
/*
		语言数据类型
*/

/*
	1.声明的局部变量必须要使用，否则就会报错，import 中的引用也是，全局变量声明后，如果不用不会报错
*/
func main() {

	/*
		字符串
	*/
	var firstName string
	var oneName, twoName string
	firstName = "Jinx"
	oneName = "Alex"
	twoName = "Bob"

	fmt.Println(firstName)
	fmt.Println(oneName,twoName)

	/*
		布尔型
	*/
	var demoBool bool = true
	fmt.Println(demoBool)

	/*
		数字类型
	*/
	//数字类型
	var demoInt int = 1
	fmt.Println(demoInt)

	//无符号 8 位整型 (0 到 255)
	var demoUint8 uint8 = 1
	fmt.Println(demoUint8)
	//无符号 16 位整型 (0 到 65535)
	var demoUint16 uint16 = 2
	fmt.Println(demoUint16)

	//无符号 32 位整型 (0 到 4294967295)
	var demoUint32 uint32 = 3
	fmt.Println(demoUint32)

	//无符号 64 位整型 (0 到 18446744073709551615)
	var demoUint64 uint64 = 4
	fmt.Println(demoUint64)

	//有符号 8 位整型 (-128 到 127)
	var demoInt8 int8 = -1
	fmt.Println(demoInt8)

	//有符号 16 位整型 (-32768 到 32767)
	var demoInt16 int16 = -2
	fmt.Println(demoInt16)

	//有符号 32 位整型 (-2147483648 到 2147483647)
	var demoInt32 int32 = -3
	fmt.Println(demoInt32)

	//有符号 64 位整型 (-9223372036854775808 到 9223372036854775807)
	var demoInt64 int64 = -4
	fmt.Println(demoInt64)

	/*
		浮点类型
	*/

	//32位浮点类型
	var demoFloat32 float32 = 1.1
	fmt.Println(demoFloat32)

	//64位浮点类型
	var demoFloat64 float64 = 1.21
	fmt.Println(demoFloat64)

	//32 位实数和虚数
	var demoComplex64 complex64 = 1+2i
	fmt.Println(demoComplex64)

	//64 位实数和虚数
	var demoComplex128 complex128 = 2+3i
	fmt.Println(demoComplex128)

	/*
		其他数字类型
	*/

	//类似 uint8
	var demoByte byte = 12
	fmt.Println(demoByte)

	//类似 int32
	var demoRune rune = 12
	fmt.Println(demoRune)

	//32 或 64 位
	var demoUint uint = 23
	fmt.Println(demoUint)

	//无符号整型，用于存放一个指针
	var demoUintptr uintptr = 0
	fmt.Println(demoUintptr)
}
