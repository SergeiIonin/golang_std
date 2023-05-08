package main

import (
	"fmt"
	"time"
)

func worker(tasks <-chan int, results chan<- int, id int) {
	for num := range tasks {
		fmt.Printf("Worker %v is executing a task for %v...\n", id, num)
		time.Sleep(1000 * time.Millisecond)
		results <- num * num
	}
}

func main() {

	tasks := make(chan int, 10)
	results := make(chan int, 10)

	for i := 0; i < 3; i++ {
		go worker(tasks, results, i)
	}

	for i := 0; i < 5; i++ {
		tasks <- i * 2
	}

	//close(tasks) // closing isn't necessary

	for i := 0; i < 5; i++ {
		res := <-results
		fmt.Println("res = ", res)
	}

}
