package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nvrmndmnm/numboozle/internal/http-server/handlers/game"
	"github.com/nvrmndmnm/numboozle/internal/http-server/handlers/pages"
)



func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	page := pages.NewPage()
	e.Renderer = pages.NewTemplate()

	e.Static("/css", "web/static/css")
	e.Static("/js", "web/static/js")

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", page)
	})

	e.GET("/game", game.GameHandler)
	e.POST("/game/click", game.ClickHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
