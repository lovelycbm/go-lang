package main

import (
	"fmt"
	"strings"
)

// naked return return 부분에 이미 명시함.
func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// func repeatMe(word ...string) {
// 	fmt.Println(word)
// }

func main() {
	test := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(test)
	//repeatMe("123", "dpjsadf", "alkjsdf", "test")
	// totalLength, uppercase := lenAndUpper(" bong mun choi")
	// fmt.Println(totalLength, uppercase)
}
