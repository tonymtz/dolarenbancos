package v1

import (
	"github.com/labstack/echo"
	"github.com/tonymtz/dolarenbancos/api/v1/banks"
	"github.com/tonymtz/dolarenbancos/api/v1/prices"
	pricesService "github.com/tonymtz/dolarenbancos/services/prices"
)

func Router(e *echo.Echo) {
	api := e.Group("/v1")

	myPricesService := pricesService.NewService()
	myPricesRoute := prices.New(myPricesService)
	myBanksRoute := banks.New()

	api.GET("/prices", myPricesRoute.Get)
	api.GET("/banks", myBanksRoute.Get)
}
