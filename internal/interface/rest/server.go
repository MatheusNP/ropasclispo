package rest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type Controllers struct {
	PlayController PlayController
}

var validate *validator.Validate

func BindInputData(ctx context.Context, c echo.Context, data interface{}) error {
	if err := c.Bind(data); err != nil {
		return HandleError(http.StatusBadRequest, fmt.Sprintf("there is a bind error on your content: %s", err.Error()))
	}

	if err := validate.Struct(data); err != nil {
		return HandleError(http.StatusBadRequest, fmt.Sprintf("there is a bind error on your content: %s", err.Error()))
	}

	return nil
}

func (c *Controllers) NewServer() {
	e := echo.New()
	validate = validator.New(validator.WithRequiredStructEnabled())

	e.GET("/hello", helloWord)
	e.POST("/play", c.PlayController.Fight)

	e.Logger.Fatal(e.Start(":1323"))
}

func helloWord(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "hello word",
	})
}
