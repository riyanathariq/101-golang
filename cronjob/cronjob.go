package cronjob

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"log"
	"time"
)

func StartCronJob() {
	sh := gocron.NewScheduler(time.Local)

	sh.Every(1).Seconds()

	_, err := sh.Do(func() {
		fmt.Println("Cronjob starting at ", time.Now().Format("15:04:05"))
	})
	if err != nil {
		log.Fatalf("Cronjob failed to start: %v", err)
		return
	}

	sh.StartAsync()

	select {}
}
