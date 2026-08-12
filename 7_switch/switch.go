package main

import "fmt"

func main() {
	// simple switch

	i := 2

	// no need to write the break here in the switch in the go
	// because it manages internally that thing
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	default:
		fmt.Println("Get lost")
	}
}
