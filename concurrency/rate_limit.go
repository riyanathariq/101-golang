package concurrency

import (
	"fmt"
	"time"
)

func RateLimiter() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(1 * time.Second)

	for req := range requests {
		<-limiter
		fmt.Println("Processing request", req, "at", time.Now().Format("15:04:05"))
	}

	fmt.Println("✅ All requests processed.")
}
