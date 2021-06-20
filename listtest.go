package main

import (
	"fmt"
	"regexp"
)

// var str list.List

func main() {
	/*这个是list 我还没有研究清楚*/
	// l := list.New()
	// l.PushBack("10.236.72.0-10.236.73.254")
	// l.PushBack("10.236.24.0-10.236.25.254")
	// l.PushFront("10.236.32.0-10.236.32.254")
	// fmt.Printf("%s", l)
	// for i := l.Front(); i != nil; i = i.Next() {
	// 	fmt.Println(i.Value)

	// }

	/*这个是正则匹配*/
	str := "a:bb ,c:dd20.5 [10.236.2.0-10.236.3.0,10.236.5.0-10.236.6.0]"
	relist := regexp.MustCompile("[0-9]+\\.[0-9]+\\.[0-9]+\\.[0-9]-[0-9]+\\.[0-9]+\\.[0-9]+\\.[0-9]")
	match := relist.FindAllString(str, -1)
	fmt.Println(match[1])

	// a:=list(str)
	// fmt.Println(l.Front().Value)
	// list2=[1,2,3,4,"sss"]

}
