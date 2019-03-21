package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	count int
	wg    sync.WaitGroup
	mutex sync.Mutex
)

func main() {

	wg.Add(2)

	go incCount(1)
	go incCount(2)

	wg.Wait()
	fmt.Printf("Final counter %d\n", count)

}

func incCount(id int) {
	defer wg.Done()

	for i := 0; i < 3; i++ {
		mutex.Lock()
		{
			value := count

			runtime.Gosched()
			value++
			count = value

		}
		mutex.Unlock()
	}
}
