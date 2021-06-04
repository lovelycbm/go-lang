package main

import (
	"fmt"
	"strings"
)

// naked return return 부분에 이미 명시함.
func lenAndUpper(name string) (lenght int, uppercase string) {
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	totalLength, uppercase := lenAndUpper(" bong mun choi")
	fmt.Println(totalLength, uppercase)
}
