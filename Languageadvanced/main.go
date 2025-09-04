package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func addTen(num *int) {
	*num += 10
}

func bet(slice *[]int) {
	for i := range *slice {
		(*slice)[i] *= 2
	}
}

func printa(wg *sync.WaitGroup) {
	start := time.Now()
	defer wg.Done()
	for i := 1; i < 10; i += 2 {
		fmt.Printf("奇数: %d\n", i)
	}
	time.Sleep(1000)
	passtime := time.Since(start)
	fmt.Println("goroutine题目2执行时间a", passtime)
}

func printb(wg *sync.WaitGroup) {
	start := time.Now()
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		fmt.Printf("偶数: %d\n", i)
	}
	time.Sleep(1000)
	passtime := time.Since(start)
	fmt.Println("goroutine题目2执行时间b", passtime)
}

func main() {
	value := 5
	fmt.Println("指针题目1：调用前值：", value)
	addTen(&value)
	fmt.Println("指针题目1：调用后的值", value)
	numbers := []int{1, 2, 3, 4, 5}
	bet(&numbers)
	fmt.Println("指针题目2", numbers)

	var wg sync.WaitGroup
	runtime.GOMAXPROCS(2)
	wg.Add(2)

	go printa(&wg)

	go printb(&wg)

	wg.Wait()

	fmt.Println("goroutine题目1：所有进程已完成")
}
