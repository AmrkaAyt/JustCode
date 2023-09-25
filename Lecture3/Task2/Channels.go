package Task2

import "sync"

func mergeChannels(channels ...chan int) chan int {
	var wg sync.WaitGroup
	merged := make(chan int, 10)

	copy := func(ch chan int) {
		defer wg.Done()
		for val := range ch {
			merged <- val
		}
	}

	wg.Add(len(channels))

	for _, ch := range channels {
		go copy(ch)
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}
