package concurrency

import (
	"fmt"
	"sync"
)

func Mutex() {
	var mu sync.Mutex
	counter := 0
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			mu.Lock()
			defer mu.Unlock()
			counter++
			fmt.Println("counter:", counter)
		}()
	}

	wg.Wait()
}
