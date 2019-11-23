package prices

import (
	"github.com/tonymtz/dolarenbancos/models"
	"github.com/tonymtz/dolarenbancos/services/prices/providers"
	"github.com/tonymtz/dolarenbancos/services/prices/providers/banamex"
	"github.com/tonymtz/dolarenbancos/services/prices/providers/banxico"
	"github.com/tonymtz/dolarenbancos/services/prices/providers/inbursa"
	"github.com/tonymtz/dolarenbancos/services/prices/providers/santander"
	"log"
	"sync"
)

var banxicoFetcher providers.Fetcher
var banamexFetcher providers.Fetcher
var santanderFetcher providers.Fetcher
var inbursaFetcher providers.Fetcher

func init() {
	banxicoFetcher = banxico.New()
	banamexFetcher = banamex.New()
	santanderFetcher = santander.New()
	inbursaFetcher = inbursa.New()
}

type Service interface {
	FetchAll() []*models.Price
}

type pricesService struct {
	Service
}

func (this *pricesService) FetchAll() []*models.Price {
	var prices []*models.Price

	functions := []providers.Fetcher{
		banxicoFetcher,
		banamexFetcher,
		santanderFetcher,
		inbursaFetcher,
	}

	channel := make(chan *models.Price)
	var waitGroup sync.WaitGroup

	for _, fetcher := range functions {
		waitGroup.Add(1)
		go getPrice(fetcher, channel, &waitGroup)
	}

	for range functions {
		prices = append(prices, <-channel)
	}

	waitGroup.Wait()

	return prices
}

func getPrice(
	fetcher providers.Fetcher,
	channel chan *models.Price,
	waitGroup *sync.WaitGroup,
) {
	defer waitGroup.Done()

	price, err := fetcher.Fetch()

	if err != nil {
		log.Fatalln(err)
	}

	channel <- price
	return
}

func NewService() Service {
	return &pricesService{}
}
