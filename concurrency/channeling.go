package concurrency

import (
	"fmt"
)

func Channeling() {
	ch := make(chan string)
	for i := 0; i <= 10; i++ {
		go func(ch chan<- string) {
			ch <- fmt.Sprintf("Repetisi channeling dengan mutex ke-%d", i)
		}(ch)
	}

	for i := 1; i <= 3; i++ {
		msg := <-ch
		fmt.Println(msg)
	}
}
