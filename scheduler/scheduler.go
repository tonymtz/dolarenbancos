package scheduler

import (
	"database/sql"
	"github.com/robfig/cron"
	pricesRepository "github.com/tonymtz/dolarenbancos/repositories/prices"
	"github.com/tonymtz/dolarenbancos/server/logger"
	"github.com/tonymtz/dolarenbancos/services/prices"
)

var pricesService prices.Service
var myPricesRepository pricesRepository.PricesRepository

func init() {
	pricesService = prices.NewService()
}

func StartScheduler(db *sql.DB) {
	myPricesRepository = pricesRepository.New(db)

	// second | minute | hour | day-of-month | month | day of week
	// https://crontab.guru/#0_*/1_*_*_*
	c := cron.New(cron.WithSeconds())
	c.AddFunc("0 */1 * * *", digest)
	c.Start()
}

func digest() {
	println("[scheduler] CRON started")

	unsavedPrices := pricesService.FetchAll()
	for _, unsavedPrice := range unsavedPrices {
		_, err := myPricesRepository.Create(unsavedPrice)

		if err != nil {
			logger.Error(err)
		}
	}

	println("[scheduler] CRON finished")
}
