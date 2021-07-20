package main

import (
	"fmt"

	"example.com/hello/mydict"
)

func main() {

	dictionary := mydict.Dictionary{}
	baseWord := "hello"

	dictionary.Add(baseWord, "First")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)

	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// hello, _ := dictionary.Search(word)
	// fmt.Println("found : ", word, "definition : ", hello)

	// err2 := dictionary.Add(word, definition)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }
	// definition, err := dictionary.Search("first")
	// // 간단한데 강력함.
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(definition)
	// }

}
