package score

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type ScoreSaver interface {
	SaveScore(ctx context.Context, userId int, score time.Duration) (int64, error)
}

func New(scoreSaver ScoreSaver) echo.HandlerFunc {
	return func(c echo.Context) error {
		number := c.FormValue("number")
		if number == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "missing number value",
			})
		}

		gameTime := c.FormValue("time")
		fmt.Println(gameTime)
		if gameTime == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "missing time value",
			})
		}

		score, err := strconv.ParseInt(gameTime, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "invalid number",
			})
		}

		if number == "25" {
			_, err = scoreSaver.SaveScore(c.Request().Context(), 1, time.Duration(score))
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{
					"error": "failed to save num" + err.Error(),
				})
			}
		}

		return nil
	}
}
