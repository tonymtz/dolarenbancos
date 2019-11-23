package banamex

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

const BANAMEX_URL = "https://finanzasenlinea.infosel.com/banamex/WSFeedJSON/service.asmx/DivisasLast?callback="

type banamex struct {
	providers.Fetcher
}

func (this *banamex) Fetch() (*models.Price, error) {
	resp, err := http.Get(BANAMEX_URL)

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
		if item.CveInstrumento == "MXNUS" {
			sell, _ := strconv.ParseFloat(item.ValorActualVenta, 64)
			buy, _ := strconv.ParseFloat(item.ValorActualCompra, 64)

			return &models.Price{
				Sell: sell,
				Buy:  buy,
				Bank: banks.Banamex.Id,
			}, nil
		}
	}

	return nil, errors.New("payload wrongly formatted")
}

func New() *banamex {
	return &banamex{}
}
