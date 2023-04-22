package di

import (
	"github.com/kondroid00/sample-server-2022/main/domain/repository"
	"github.com/kondroid00/sample-server-2022/main/domain/service"
	"github.com/kondroid00/sample-server-2022/main/infra"
	"github.com/kondroid00/sample-server-2022/main/server"
	"github.com/kondroid00/sample-server-2022/main/server/controller"
	"github.com/kondroid00/sample-server-2022/main/usecase"
	"github.com/kondroid00/sample-server-2022/package/errors"
	"go.uber.org/dig"
)

func NewHandler(i infra.Service) (server.Handler, error) {
	c := dig.New()

	if err := provide(c, i); err != nil {
		return nil, err
	}

	if err := invoke(c); err != nil {
		return nil, err
	}

	var h server.Handler
	if err := c.Invoke(func(handler server.Handler) {
		h = handler
	}); err != nil {
		return nil, err
	}

	return h, nil
}

func provide(c *dig.Container, i infra.Service) error {
	mError := errors.NewMultiError()

	// infra
	mError.Append(c.Provide(func() infra.Service {
		return i
	}))

	// handler
	mError.Append(c.Provide(server.NewHandler))

	// controller
	mError.Append(c.Provide(controller.NewUser))

	// usecase
	mError.Append(c.Provide(usecase.NewUser))

	// domain service
	mError.Append(c.Provide(service.NewUser))

	// domain repository
	mError.Append(c.Provide(repository.NewUser))

	return mError.GetError()
}

func invoke(c *dig.Container) error {
	mError := errors.NewMultiError()

	// handler
	mError.Append(c.Invoke(server.NewHandler))

	// controller
	mError.Append(c.Invoke(controller.NewUser))

	// usecase
	mError.Append(c.Invoke(usecase.NewUser))

	// domain service
	mError.Append(c.Invoke(service.NewUser))

	// domain repository
	mError.Append(c.Invoke(repository.NewUser))

	return mError.GetError()
}
