package main

import (
	"fmt"
)

// maps -> hash, object, dict
func main() {
	// creating map

	// m := make(map[string]string)

	// setting an element
	// m["name"] = "golang"
	// m["area"] = "backend"

	// get an element
	// fmt.Println(m["name"], m["area"])
	// IMP: if key does not exists in the map then it returns no value

	// m := make(map[string]int)
	// m["age"] = 30
	// m["price"] = 50
	// fmt.Println(m["age"])
	// fmt.Println(len(m))

	// delete(m, "prce") // this will not throw the error but likewise here it will not delete and all
	// clear(m)

	// fmt.Println(m)
	// fmt.Println(len(m))
	// // fmt.Println(m)

	m := map[int]int{4: 40, 56: 3}

	v, ok := m[4] // here the ok is key and v is the value so checking the
	fmt.Println(v)
	if ok {
		fmt.Println("all ok")
	} else {
		fmt.Println("not ok")
	}

	// m1 := map[string]int{"price": 40, "phones": 3}
	// m2 := map[string]int{"price": 40, "phones": 8}
	// fmt.Println(maps.Equal(m1, m2))

}
