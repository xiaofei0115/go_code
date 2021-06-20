package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company  string
	Subjects []string `json:list.List`
}

func main() {
	jsonbuff := `{
		"company":"zhczGO",
		"subjects":[
			"go",
			"python",
			"Test"
		]

	}`
	var temp IT
	err := json.Unmarshal([]byte(jsonbuff), &temp)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	a := temp.Subjects
	fmt.Printf(a)

	// fmt.Printf("tmp=%+v\n", temp)

}
