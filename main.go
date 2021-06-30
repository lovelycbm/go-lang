package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"steak", "ramen"}
	choi := person{name: "choi", age: 35, favFood: favFood}
	fmt.Println(choi)

}
