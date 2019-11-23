package banxico

import (
	"encoding/json"
	"github.com/tonymtz/dolarenbancos/models"
	"github.com/tonymtz/dolarenbancos/repositories/banks"
	"github.com/tonymtz/dolarenbancos/services/prices/providers"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const BANXICO_URL = "https://www.banxico.org.mx/SieAPIRest/service/v1/series/SF43718/datos/oportuno"
const TOKEN_BMX = "be79a3f8aa6ef6b7d092c9b24cd8c21457bd8fe088ed430a3b3effbb6a8486eb"

type banxico struct {
	providers.Fetcher
}

func (this *banxico) Fetch() (*models.Price, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", BANXICO_URL, nil)

	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("Bmx-Token", TOKEN_BMX)
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var response Response
	if err = json.Unmarshal(body, &response); err != nil {
		log.Fatalln(err)
	}

	data := response.Bmx.Series[0].Datos[0]
	price, err := strconv.ParseFloat(data.Dato, 64)

	if err != nil {
		log.Fatalln(err)
	}

	return &models.Price{
		Sell: price,
		Buy:  price,
		Bank: banks.Banxico.Id,
	}, nil
}

func New() *banxico {
	return &banxico{}
}
