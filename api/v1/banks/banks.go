package banks

import (
	"github.com/labstack/echo"
	"github.com/tonymtz/dolarenbancos/repositories/banks"
	"net/http"
)

type BanksRoute interface {
	Get(echo.Context) error
}

type banksRoute struct {
	BanksRoute
}

func (this *banksRoute) Get(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, banks.Banks)
}

func New() *banksRoute {
	return &banksRoute{}
}
