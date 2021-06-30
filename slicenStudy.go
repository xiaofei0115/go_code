package main

import "fmt"
//这个有点像list的操作，使用数据的方式
//字符串能否转换为数组了？
func main() {
	arry := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arry)
	// s1 := arry[2:6]
	// s2 := s1[3:5]
	// fmt.Println(s2)
	arry2 := [...]string{"a", "b", "1"}


	
	for s, a := range arry2 {
		fmt.Println(s, a)
	}
	fmt.Println("arry2=" + arry2[1])
	fmt.Println(arry2[1:])
}
