package prices

import (
	"github.com/labstack/echo"
	"github.com/tonymtz/dolarenbancos/services/prices"
	"net/http"
)

type PricesRoute interface {
	Get(echo.Context) error
}

type pricesRoute struct {
	PricesRoute
	pricesService prices.Service
}

func (this *pricesRoute) Get(ctx echo.Context) error {
	result := this.pricesService.FetchAll()
	return ctx.JSON(http.StatusOK, result)
}

func New(pricesService prices.Service) *pricesRoute {
	return &pricesRoute{pricesService: pricesService}
}
