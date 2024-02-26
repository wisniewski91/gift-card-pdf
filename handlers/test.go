package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
	test_view "github.com/wisniewski91/gift-card-pdf/views/test"
)

func TestHandler(c echo.Context) error {

	return render(c, test_view.Render())

}
func TestPostHandler(c echo.Context) error {

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error reading request body")
	}

	// Print the request body
	fmt.Println("Request Body:", string(body))

	// Optionally, you can also send a response back to the client
	return c.String(http.StatusOK, string(body))

}
