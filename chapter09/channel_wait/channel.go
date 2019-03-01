package main

import "fmt"
import "time"

func printNumber2(ch chan bool) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d ", i)
	}
	ch <- true
}

func printLetters2(ch chan bool) {
	for i := 'A'; i < 'A' + 10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c ", i)
	}
	ch <- true
}

func main() {
	ch1, ch2 := make(chan bool), make(chan bool)
	go printNumber2(ch1)
	go printLetters2(ch2)
	<- ch1
	<- ch2
}
