package main

import "fmt"

func main() {
	var name string = "golang"

	// here the warn varibale infers the datatype would be like teh string only liekwise

	var warn = "hello"
	var num = 1

	fmt.Println(warn)

	fmt.Println(name)

	fmt.Println(num)

	// shorthand syntax
	age := 21

	fmt.Println(age)
}
