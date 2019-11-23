package prices

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/tonymtz/dolarenbancos/models"
)

type PricesRepository interface {
	Create(price *models.Price) (*models.Price, error)
}

type pricesRepository struct {
	PricesRepository
	database *sql.DB
}

func (this *pricesRepository) Create(newPrice *models.Price) (*models.Price, error) {
	var lastInsertedId int64

	err := this.database.QueryRow(
		"INSERT INTO prices (sell, buy, bank) VALUES ($1, $2,$3) RETURNING id",
		newPrice.Sell,
		newPrice.Buy,
		newPrice.Bank,
	).Scan(&lastInsertedId)

	if err != nil {
		return nil, err
	}

	newPrice.Id = lastInsertedId

	return newPrice, nil
}

func New(database *sql.DB) PricesRepository {
	return &pricesRepository{database: database}
}
