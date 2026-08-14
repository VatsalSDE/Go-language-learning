package main

import (
	"fmt"
	"sync"
)

// func task(id int, w *sync.WaitGroup) {
// 	defer w.Done()
// 	fmt.Println("doing task", id)
// }

func task(id int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("doing task", id)
}

func main() {
	var wg sync.WaitGroup

	// for i := 0; i <= 10; i++ {
 
	// 	wg.Add(1)
	// 	go task(i, &wg)
	// }

	for i := 0; i <= 10; i++ {
		wg.Add(2)
		go task(i, &wg)
		go task(i, &wg)
	}

	// time.Sleep(time.Millisecond * 1) // wait for 100 milliseconds to let all goroutines finish
	wg.Wait()
}
