package santander

import (
	"encoding/json"
	"errors"
	"github.com/tonymtz/dolarenbancos/models"
	"github.com/tonymtz/dolarenbancos/repositories/banks"
	"github.com/tonymtz/dolarenbancos/services/prices/providers"
	"io/ioutil"
	"net/http"
	"strconv"
)

const SANTANDER_URL = "https://finanzasenlinea.infosel.com/Santanderfeed/Feed.asmx/Divisas"

type santander struct {
	providers.Fetcher
}

func (this *santander) Fetch() (*models.Price, error) {
	resp, err := http.Get(SANTANDER_URL)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var response []Intrumento
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	for _, item := range response {
		if item.Instrumento == "MXNUS" {
			sell, _ := strconv.ParseFloat(item.Venta, 64)
			buy, _ := strconv.ParseFloat(item.Compra, 64)

			return &models.Price{
				Sell: sell,
				Buy:  buy,
				Bank: banks.Santander.Id,
			}, nil
		}
	}

	return nil, errors.New("payload wrongly formatted")
}

func New() *santander {
	return &santander{}
}
