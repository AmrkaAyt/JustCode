package Task1

import "sync"

func SimpleDeadlock() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)

	go func() {
		defer wg.Done()
		val := <-ch
		_ = val
	}()

	wg.Wait()
}

func DeadlockWithoutBuffer() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)

	go func() {
		defer wg.Done()
		ch <- 42
	}()

	wg.Wait()
}
