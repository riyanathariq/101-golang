package main

import (
	"fmt"
	"github.com/riyanathariq/101-golang/concurrency"
	"github.com/riyanathariq/101-golang/cronjob"
	"github.com/riyanathariq/101-golang/redirect"
	"os"
)

// Used to refer the submodel on package
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [topic]")
		return
	}

	switch os.Args[1] {
	case "rate_limit":
		concurrency.RateLimiter()
	case "cronjob":
		cronjob.StartCronJob()
	case "channeling":
		concurrency.Channeling()
	case "mutex":
		concurrency.Mutex()
	case "redirect":
		redirect.Redirect()
	default:
		fmt.Println("Unknown model")
	}
}
