package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"radityaarya/go-clean-boilerplate/common/utils"
)

type ArticlesHandler struct {
}

func (h *ArticlesHandler)GetListArticles(c echo.Context) error {
	return c.JSON(http.StatusOK, utils.ResponseSuccess())
}