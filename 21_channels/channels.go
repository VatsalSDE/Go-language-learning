// package main

// import (
// 	"fmt"
// 	"time"
// )

// // send
// // func processNum(numChan chan int) {
// // 	for num := range numChan {
// // 		fmt.Println("processing number", num)
// // 		time.Sleep(time.Second)
// // 	}

// // }

// // receive
// // func sum(result chan int, num1 int, num2 int) {
// // 	numResult := num1 + num2
// // 	result <- numResult
// // }

// // goroutine synchronizer
// // func task(done chan bool) {
// // 	defer func() { done <- true }()

// // 	fmt.Println("processing...")
// // }

// func emailSender(emailChan <-chan string, done chan<- bool) {
// 	defer func() { done <- true }()

// 	for email := range emailChan {
// 		fmt.Println("sending email to", email)
// 		time.Sleep(time.Second)
// 	}
// }

// func main() {
// 	// chan1 := make(chan int)
// 	// chan2 := make(chan string)

// 	// go func() {
// 	// 	chan1 <- 10
// 	// }()

// 	// go func() {
// 	// 	chan2 <- "pong"
// 	// }()

// 	// for i := 0; i < 2; i++ {
// 	// 	select {
// 	// 	case chan1Val := <-chan1:
// 	// 		fmt.Println("received data from chan1", chan1Val)
// 	// 	case chan2Val := <-chan2:
// 	// 		fmt.Println("received data from chan2", chan2Val)
// 	// 	}
// 	// }

// 	emailChan := make(chan string, 100)
// 	// done := make(chan bool)

// 	// go emailSender(emailChan, done)

// 	// for i := 0; i < 5; i++ {
// 	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
// 	// }

// 	// fmt.Println("done sending.")

// 	// this is important
// 	// close(emailChan)
// 	// <-done
// 	emailChan <- "1@example.com"
// 	emailChan <- "2@example.com"

// 	fmt.Println(<-emailChan)
// 	fmt.Println(<-emailChan)

// 	// done := make(chan bool)
// 	// go task(done)

// 	// <-done // block

// 	// result := make(chan int)
// 	// go sum(result, 4, 5)
// 	// res := <-result // blocking

// 	// fmt.Println(res)
// 	// numChan := make(chan int)

// 	// go processNum(numChan)

// 	// for {
// 	// 	numChan <- rand.Intn(100)
// 	// }

// 	// messageChan := make(chan string)

// 	// messageChan <- "ping" // blocking

// 	// msg := <-messageChan

// 	// fmt.Println(msg)

// }

// // func processNum(numchan chan int) {
// // 	msg := <-numchan

// // 	fmt.Println(msg)
// // }

// // // this is the example of sending that is from the main
// // func processNum(numchan chan int) {

// // 	for num := range numchan {

// // 		fmt.Println(num)
// // 		time.Sleep(time.Second)

// // 	}
// // }

// // func task(done chan bool) {

// // 	// this is a cleanup function and it will like wise run when the whole thing completes or not likewise
// // 	defer func() {
// // 		done <- true
// // 	}()

// // 	fmt.Println("Process is running ")

// // }

// // func sum(result chan int, num1 int, num2 int) {
// // 	numResult := num1 + num2
// // 	result <- numResult
// // }

// // func main() {
// // 	// messageChan := make(chan string) // channel created
// // 	numchan := make(chan int)

// // 	go processNum(numchan)

// // 	numchan <- 1

// // 	// for {
// // 	// 	numchan <- rand.Intn(100)
// // 	// }

// // 	result := make(chan int)

// // 	go sum(result, 2, 3)

// // 	res := <-result

// // 	fmt.Println(res)
// // 	// messageChan <- "ping" // we are sending the data to the message chan this is the blocking operation

// // 	// msg := <-messageChan

// // 	// fmt.Println(msg)
// // }

package main

import "fmt"

func worker(jobs <-chan int, results chan<- int) {
	for job := range jobs {
		result := job * 2
		results <- result
	}
}

func main() {
	jobs := make(chan int, 3)
	results := make(chan int, 3)

	go worker(jobs, results)

	jobs <- 10
	jobs <- 20
	jobs <- 30

	close(jobs)

	for i := 0; i < 3; i++ {
		result := <-results
		fmt.Println(result)
	}
}
