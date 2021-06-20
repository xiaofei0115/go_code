package main

import "fmt"

func main() {
	age := 18
	// &符号+变量 就可以获取内存地址
	fmt.Println(&age) //0xc000012088
	//定义一个指着变量
	//ptr 指针变量的名字 *int
	// var ptr *int =0xc000012088
	var ptr *int = &age
	fmt.Println(*ptr) //取得值

	num := 10
	fmt.Println(num)
	var ptr1 *int = &num
	*ptr1 = 20
	fmt.Println(num)
}
