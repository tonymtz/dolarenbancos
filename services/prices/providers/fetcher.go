package providers

import "github.com/tonymtz/dolarenbancos/models"

type Fetcher interface {
	Fetch() (*models.Price, error)
}
