package main

import "fmt"

func a(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}
func b(ch1 <-chan int, ch2 chan<- int) {
	for {
		num, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- num
	}

	close(ch2)
}

func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)

	go a(ch1)
	go b(ch1, ch2)

	for num := range ch2 {
		fmt.Println(num)
	}

}
