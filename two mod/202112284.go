package main

import "fmt"

/*
	集合
*/
func main() {
	var countryCapitalMap map[string]string

	/*创建集合*/
	countryCapitalMap = make(map[string]string)

	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	/*使用 key 输出 map 值*/
	for country := range countryCapitalMap {
		fmt.Println("Capital of ", country, " is ", countryCapitalMap[country])
	}
	/*查看元素是否在集合中存在*/
	captial, ok := countryCapitalMap["United States"]
	if ok {
		fmt.Println("Capital of United States is", captial)
	} else {
		fmt.Println("Capital of United States is not present")
	}

	/*
		打印并不会按顺序打印
	*/
	countryCapitalMap2 := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New Delhi"}

	fmt.Println("原始 map")

	for k, v := range countryCapitalMap2 {
		fmt.Println("Capital of ", k, " is ", v)
	}

	delete(countryCapitalMap2, "France")
	fmt.Println("Entry for France is deleted")

	fmt.Println("删除元素后map")

	for k, v := range countryCapitalMap2 {
		fmt.Println("Capital of ", k, " is ", v)
	}
}
