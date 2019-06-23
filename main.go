package main

import (
	interfaceHTTP "radityaarya/go-clean-boilerplate/articles/interface/http"
	"radityaarya/go-clean-boilerplate/common/utils"

	"github.com/labstack/echo/v4"
)

var (
	ar = interfaceHTTP.ArticlesRoute{}
)

func main() {
	e := echo.New()

	articles := e.Group("/articles")
	ar.Route(articles)

	e.HTTPErrorHandler = utils.OverrideErrorMethod

	e.Logger.Fatal(e.Start(":3000"))
}
