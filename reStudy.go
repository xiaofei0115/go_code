package main

import (
	"fmt"
	"regexp"
)

// "regexp"
const test = "hello goland a bb@qq.com ccdd@qq.com sss@ww.com"

func main() {
	reslist := regexp.MustCompile("[A-Za-z0-9]+@[A-Za-z0-9]+\\.[A-Za-z0-9]+")
	match := reslist.FindAllString(test, -1)
	fmt.Println(match[0])

	//数组是[a b c] []string 这个是有下角标的
	//列表是[a,b,c]
	// for i := 0; i < len(match); i++ {
	// 	fmt.Print(match[i])

	// }

	// res := regexp.MustCompile("ab")
	// reslist := regexp.MustCompile("[A-Za-z0-9]+@[A-Za-z0-9]+\\.[A-Za-z0-9]+")
	// match := reslist.FindAllString(test, -1)
	// for m := range match {
	// 	fmt.Println(m[0])

}
