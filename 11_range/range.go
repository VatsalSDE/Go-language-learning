package main

import "fmt"

// iterating over data structures

func main() {

	nums := []int{6, 7, 8}

	for _, num := range nums {
		fmt.Println(num)
	}

	n := len(nums)
	for i := 0; i < n; i++ {
		fmt.Println(nums[i])
		i++
	}

	m := map[string]string{"fname": "john", "lname": "doe"}

	for k, v := range m {
		fmt.Println(k, v)
		fmt.Println(k, v)
	}

	// for k := range m {
	// 	fmt.Println(k)
	// }

	// unicode code point rune
	// starting byte of rune
	// 300 -> 1 byte , 2 byte

	for i, c := range "golang" {
		fmt.Println(i, string(c))
	}

}
