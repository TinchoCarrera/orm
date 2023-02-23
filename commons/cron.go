package commons

import (
	"fmt"
	"log"

	"github.com/mileusna/crontab"
)

func SetCron() {
	cron := crontab.New()

	fmt.Println("Cron creado cada 1 minuto")

	err := cron.AddJob("* * * * *", func() {
		log.Println("Ejecución cada un minuto")
	})

	if err != nil {
		log.Println("Hay un error en el Cron")
	}
}
