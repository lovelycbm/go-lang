package main

import (
	"fmt"

	"example.com/hello/mydict"
)

func main() {

	dictionary := mydict.Dictionary{"first": "First word"}
	definition, err := dictionary.Search("first")
	// 간단한데 강력함.
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

}
