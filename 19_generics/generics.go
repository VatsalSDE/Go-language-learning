package main

import "fmt"

// func printSlice[T comparable, V string](items []T, name V) {
// 	for _, item := range items {
// 		fmt.Println(item, name)
// 	}
// }

func printSlice[T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func printSlicepipe[T int | string | bool](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func printStringSlice(items []string) {
	for _, item := range items {
		fmt.Println(item)
	}
}

//
// LIFO
type stack[T any] struct {
	elements []T
}

func main() {
	// myStack := stack[string]{
	// 	elements: []string{"golang"},
	// }

	// fmt.Println(myStack)

	// nums := []int{1, 2, 3}
	names := []string{"golang", "typescript"}
	names2 := []int{1, 2, 3}
	// vals := []bool{true, false, true}
	// printStringSlice(names)
	// printSlice(vals, "john")

	printSlice(names)
	printSlice(names2)
}
