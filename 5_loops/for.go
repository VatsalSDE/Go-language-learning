package main

import "fmt"

// for is the only construct in the go for the looping

func main() {
	i := true
	j := 1

	for i == false {
		fmt.Println(j)
		j++
	}

	for i := 1; i <= 10; i++ {
		fmt.Println("Hello", i)
	}

	for i := range 12 {
		fmt.Println(i)
	}
}
