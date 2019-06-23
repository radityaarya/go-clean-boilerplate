package http

import (
	"github.com/labstack/echo/v4"
)

var handler ArticlesHandler

type ArticlesRoute struct {
}

func (r *ArticlesRoute) Route(route *echo.Group) {
	route.GET("", handler.GetListArticles)
}