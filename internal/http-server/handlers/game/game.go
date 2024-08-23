package game

import (
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Game struct {
	Mode  int
	Grid  [5][5]int
	Start time.Time
	End   time.Time
}

func shuffleNumbers() [5][5]int {
	numbers := rand.Perm(25)
	var grid [5][5]int
	for i := 0; i < 25; i++ {
		grid[i/5][i%5] = numbers[i] + 1
	}
	return grid
}

func GameHandler(c echo.Context) error {
	mode := c.QueryParam("mode")
	game := Game{
		Mode: 1,
		Grid: shuffleNumbers(),
	}
	if mode == "2" {
		game.Mode = 2
	}

	tmpl := template.Must(template.ParseFiles("web/templates/game.html"))

	return tmpl.Execute(c.Response(), game)
}

func ClickHandler(c echo.Context) error {
	numberStr := c.FormValue("number")
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid number")
	}

	_ = number

	return c.NoContent(http.StatusOK)
}
