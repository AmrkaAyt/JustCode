package main

import (
	"Lecture3/Task2"
	"fmt"
)

func main() {
	/*	fmt.Println("Simple Deadlock:")
		Task1.SimpleDeadlock()

		fmt.Println("Deadlock Without Buffer:")
		Task1.DeadlockWithoutBuffer()*/
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for i := 6; i <= 10; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	go func() {
		for i := 11; i <= 15; i++ {
			ch3 <- i
		}
		close(ch3)
	}()

	merged := Task2.MergeChannels(ch1, ch2, ch3)

	for val := range merged {
		fmt.Println(val)

	}
}
