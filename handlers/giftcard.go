package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wisniewski91/gift-card-pdf/services"
	"github.com/wisniewski91/gift-card-pdf/types"
	giftcardtemplate "github.com/wisniewski91/gift-card-pdf/views/giftCardForm"
)

func GiftCardFormHandler(c echo.Context) error {
	return render(c, giftcardtemplate.Render())
}

func GiftCardPostHandler(c echo.Context) error {

	var giftCardData types.GiftCard

	err := c.Bind(&giftCardData)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	services.GiftCardGenerator(giftCardData)

	return c.String(http.StatusOK, "Great")
}
