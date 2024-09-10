package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nvrmndmnm/numboozle/internal/config"
	"github.com/nvrmndmnm/numboozle/internal/http-server/handlers/game"
	"github.com/nvrmndmnm/numboozle/internal/http-server/handlers/pages"
	"github.com/nvrmndmnm/numboozle/internal/storage"
)

func main() {
	config := config.MustLoadConfig()

	db, err := storage.InitDB(config.Driver, config.Datasource)
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	defer db.Close()

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

	e.GET("/health", func(c echo.Context) error { return c.String(http.StatusOK, "OK") })
	e.Logger.Fatal(e.Start(":8080"))
}
