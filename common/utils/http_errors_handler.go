package utils

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func OverrideErrorMethod(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if report.Code == 404 {
		c.JSON(http.StatusNotFound, ResponseNotFound())
		return
	}

	if !ok || report.Code == 500 {
		c.JSON(http.StatusInternalServerError, ResponseInternalServerError())
		return
	}
}