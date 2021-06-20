package main

import "fmt"

func main() {
	str := "hello golang"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])

	}

}
