package main

import (
	"fmt"
)

func main() {
	//数组是可以取下标的 ，arr1[0]
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	fmt.Println(arr1[0], arr2, arr3)
	for i, v := range arr3 {
		fmt.Println(i, v)
	}
}
