package main

import "fmt"

/*
	go语言范围
	Range表达式 				| 第一个值 	| 第二个值[可选的]
 	Array 或者 slice a [n]E	| 索引 i int	| a[i] E
 	String s string type	| 索引 i int	| rune int
 	map m map[K]V	 		| 键 k K	| 值 m[k] V
 	channel c chan E	 	| 元素 e E	| none


*/
func main() {
	//使用range去求一个slice的和，使用数组跟这个非常类似
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum : ", sum)
	//在数组上使用range将传入index和值两个变量，上面例子我们不需要使用该元素的序号，所以使用空白符"_"来省略，有时候我们确实需要知道它的索引。
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index : ", i)
		}
	}

	//range也可以用在map的键值对上。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	//range也可以用来枚举unitcode字符串。第一个参数是字符的索引，第二个是字符本身
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
