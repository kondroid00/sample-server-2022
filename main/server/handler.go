package server

import (
	"github.com/kondroid00/sample-server-2022/main/infra"
	"github.com/kondroid00/sample-server-2022/main/server/controller"
	"github.com/labstack/echo/v4"
	"net/http"
)

type (
	Handler interface {
		// healthcheck
		// (GET /healthcheck)
		GetHealthcheck(ctx echo.Context) error

		controller.User
	}

	HandlerImpl struct {
		infra infra.Service
		controller.User
	}
)

func NewHandler(infra infra.Service, cUser controller.User) Handler {
	return &HandlerImpl{
		infra: infra,
		User:  cUser,
	}
}

func (impl *HandlerImpl) GetHealthcheck(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "OK")
}
