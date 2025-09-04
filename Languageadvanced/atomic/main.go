package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	counter int64
	wg      sync.WaitGroup
)

func incatomic() {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		atomic.AddInt64(&counter, 1)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go incatomic()
	}
	wg.Wait()
	fmt.Println("计数器值：", atomic.LoadInt64(&counter))
}
