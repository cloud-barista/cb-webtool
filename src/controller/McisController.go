package controller

import (
	_ "fmt"
	"net/http"
	_ "net/http"

	"github.com/labstack/echo"
)

func McisList(c echo.Context) error {
	return c.Render(http.StatusOK, "MCISlist.html", nil)
}
