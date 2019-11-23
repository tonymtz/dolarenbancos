package inbursa

import (
	"github.com/gocolly/colly"
	"github.com/tonymtz/dolarenbancos/models"
	"github.com/tonymtz/dolarenbancos/repositories/banks"
	"github.com/tonymtz/dolarenbancos/services/prices/providers"
	"log"
	"strconv"
	"strings"
)

const INBURSA_URL = "https://www.inbursa.com/portal/"

type inbursa struct {
	providers.Fetcher
}

func (this *inbursa) Fetch() (*models.Price, error) {
	price := &models.Price{}

	c := colly.NewCollector()

	c.OnHTML("#Divisas tbody tr:nth-child(1)", func(e *colly.HTMLElement) {
		price.Buy = cleanInput(e.ChildText("td:nth-child(2)"))
		price.Sell = cleanInput(e.ChildText("td:nth-child(3)"))
		price.Bank = banks.Inbursa.Id
	})

	c.Visit(INBURSA_URL)

	return price, nil
}

func New() *inbursa {
	return &inbursa{}
}

func cleanInput(dirty string) float64 {
	split := strings.Split(dirty, "\n")
	trimmed := strings.TrimSpace(split[1])
	price, err := strconv.ParseFloat(trimmed, 64)

	if err != nil {
		log.Fatalln(err)
	}

	return price
}
