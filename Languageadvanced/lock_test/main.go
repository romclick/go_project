package main

import (
	"fmt"
	"sync"
)

var (
	x    int64
	wg   sync.WaitGroup
	lock sync.Mutex
)

func add() {
	for i := 0; i < 1000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()

}

func main() {

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go add()
	}
	wg.Wait()
	fmt.Println(x)
}
