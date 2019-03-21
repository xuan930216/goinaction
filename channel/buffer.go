package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4
	taskLoad         = 10
)

var wg sync.WaitGroup

func init() {
	//init the package
	rand.Seed(time.Now().Unix())
}

func main() {
	//create a buffer channel to manage the task load
	tasks := make(chan string, taskLoad)
	//Lauch goroutines to handle tasks
	wg.Add(numberGoroutines)
	for g := 1; g <= numberGoroutines; g++ {
		go worker(tasks, g)
	}
	//add a bunch of work to get done
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}
	close(tasks)
	wg.Wait()

}

func worker(tasks chan string, worker int) {
	defer wg.Done()
	for {
		task, ok := <-tasks
		if !ok {
			//this means the channel is empty or closed
			fmt.Printf("Worker: %d : Shutting Down\n", worker)
			return
		}
		//Display we are starting the task
		fmt.Printf("Worker %d : Started %s\n", worker, task)

		//randomly wait to simulate work time
		sleep := rand.Intn(10)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		//display we finish the work
		fmt.Printf("Worker %d : Completed %s\n", worker, task)
	}
}
