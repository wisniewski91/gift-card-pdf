package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/wisniewski91/gift-card-pdf/services"
)

func PlaceOrderHandler(c echo.Context) error {
	services.Order()
	return nil
}
