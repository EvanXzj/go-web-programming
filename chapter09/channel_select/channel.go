package main

import (
	"fmt"
	// "time"
)

func callerA(c chan string) {
	c <- "Hello from A"
	close(c)
}

func callerB(c chan string) {
	c <- "Hello from B"
	close(c)
}

func main() {
	a, b := make(chan string), make(chan string)
	go callerA(a)
	go callerB(b)
	var msg string
	openA, openB := true, true
	for openA || openB {
		select {
		case msg, openA = <- a:
			if openA {
				fmt.Println(msg)
			}
		case msg, openB = <- b:
			if openB {
				fmt.Println(msg)
			}
		}
	}
}

// func main() {
// 	msg1, msg2 := "", ""
// 	a, b := make(chan string), make(chan string)
// 	go callerA(a)
// 	go callerB(b)
// 	for i:=0; i < 5; i++ {
// 		time.Sleep(1 * time.Microsecond)
// 		select {
// 		case msg1 = <- a:
// 			fmt.Println(msg1)
// 		case msg2 = <- b:
// 			fmt.Println(msg2)
// 		default:
// 			fmt.Println("default")
// 		}
// 	}
// }
