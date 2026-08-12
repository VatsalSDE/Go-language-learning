package main

import "fmt"

func main() {

	for i := 10; i <= 50; i++ {

		if i%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}
}
