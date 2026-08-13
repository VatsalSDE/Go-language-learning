package main

import "fmt"

func counter() func() int {

	count := 0

	// here this is the closure and it will likewise keep the variable declared outside now likewise

	return func() int {
		count += 1
		return count
	}

}

func main() {
	// so here it is returning the function so fiorst have to cathc it int functions likewise

	increment := counter()
	increment2 := counter()

	fmt.Println(increment())
	fmt.Println(increment())

	fmt.Println(increment2) // so here only the fucntion address will be likewise printed
	fmt.Println(increment2)
}
