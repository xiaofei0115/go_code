package main

import (
	"fmt"
)

//学习switch
func main() {
	var score int = 100
	switch score / 10 {
	case 10:
		fmt.Println("您的登记为A")
	}

}
