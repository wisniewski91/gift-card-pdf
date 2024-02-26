package server

import (
	"github.com/labstack/echo/v4"
	"github.com/wisniewski91/gift-card-pdf/handlers"
)

func StartServer() {
	e := echo.New()

	e.Static("/", "static/bg")
	e.Static("/assets", "static/assets")

	e.GET("/test", handlers.TestHandler)
	e.POST("/test", handlers.TestPostHandler)
	e.GET("/giftcard", handlers.GiftCardFormHandler)
	e.POST("/order", handlers.PlaceOrderHandler)

	e.Logger.Fatal(e.Start(":8005"))
}
