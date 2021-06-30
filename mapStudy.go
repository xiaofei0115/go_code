package main

import (
	"fmt"
)

// json转map函数，通用
// func JSONToMap(str string) map[string]interface{} {

// 	var tempMap map[string]interface{}

// 	err := json.Unmarshal([]byte(str), &tempMap)

// 	if err != nil {
// 		panic(err)
// 	}

// 	return tempMap
// }

func main() {
	// s := `{
	// 	"name":"jqw",
	// 	"age":"[10.236.3.0-10.236.3.254,10.236.7.0-10.236.8.0]"
	// }`
	// mapa := JSONToMap(s)
	// ip := mapa["age"]
	// fmt.Print(ip)

	//map的声明方法： map[string]
	m := map[string]string{
		"name": "tufei",
		"age":  "12",
		"sex":  "f",
	}
	m2 := make(map[string]int)
	var m3 map[string]int
	//遍历的方法：
	for k, v := range m {
		fmt.Println(k, v)

	}
	//如何判断是否有key：

	if nameA, ok := m["name`"]; ok {
		fmt.Printf(nameA)
	} else {
		fmt.Println("you wenti")
	}
	// 增加元素
	m2["ss"] = 16
	// 删除map元素
	delete(m, "name")
	fmt.Println(m, m2, m3)

}
