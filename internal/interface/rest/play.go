package rest

import (
	"net/http"

	"github.com/MatheusNP/ropasclispo/internal/domain"
	"github.com/MatheusNP/ropasclispo/internal/service"
	"github.com/labstack/echo"
)

type PlayController interface {
	Fight(c echo.Context) error
}

type play struct {
	playService service.Play
}

func NewPlayController(playService service.Play) PlayController {
	return &play{
		playService: playService,
	}
}

func (p play) Fight(c echo.Context) error {
	ctx := c.Request().Context()

	var payload domain.FightRequest

	if err := BindInputData(ctx, c, &payload); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	resp, _ := p.playService.Result(ctx, payload)

	return c.JSON(http.StatusOK, resp)
}
