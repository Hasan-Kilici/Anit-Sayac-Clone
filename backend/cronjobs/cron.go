package cronjobs

import (
	"anitsayac/scrapper"
	"github.com/robfig/cron/v3"
	"log"
)

func InitializeCron() *cron.Cron {
	c := cron.New()

	_, err := c.AddFunc("0 0 * * *", func() {
		log.Println("Running scheduled ScrapeData task...")
		scrapper.ScrapeData()
	})

	if err != nil {
		log.Fatalf("Error scheduling cron job: %v", err)
	}

	c.Start()

	return c
}
