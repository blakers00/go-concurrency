package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 1)
	go func() {
		defer close(ch)
		//time.Sleep(2 * time.Second)
		ch <- "passing time version 1"
	}()

	select {
	case v := <-ch:
		fmt.Printf("case 1: %s\n", v)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout version 1")
	default:
		fmt.Println("no message received")
	}

	ch2 := make(chan string, 1)
	go func() <-chan string {
		defer close(ch2)
		time.Sleep(1 * time.Second)
		ch2 <- "Passing time version 2\n"
		return ch2
	}()

	for i := 0; i <= cap(ch2); i++ {
		select {
		case k := <-ch2:
			fmt.Printf("case 2: %s", k)
		case <-time.After(2 * time.Second):
			fmt.Println("timeout version 2 ")
		default:
			continue

		}
	}
}
