package main

import (
	"fmt"
	"time"
)

func thrower(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("Threw >> ", i)
	}
}

func catcher(ch chan int) {
	for i := 0; i < 5; i++ {
		num := <-ch
		fmt.Println("Caught <<", num)
	}
}

func main() {
	c := make(chan int, 3)
	go thrower(c)
	go catcher(c)
	time.Sleep(100 * time.Millisecond)
}
