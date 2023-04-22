package controller

import (
	"github.com/kondroid00/sample-server-2022/main/domain/ent"
	"github.com/kondroid00/sample-server-2022/main/domain/ent/schema"
	"github.com/kondroid00/sample-server-2022/main/infra"
	cEnt "github.com/kondroid00/sample-server-2022/main/interface/converter/ent"
	"github.com/kondroid00/sample-server-2022/main/interface/converter/response"
	"github.com/kondroid00/sample-server-2022/main/interface/openapi"
	"github.com/kondroid00/sample-server-2022/main/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type (
	User interface {
		// Create New User
		// (POST /user)
		PostUser(ctx echo.Context) error
		// Get User Info by User ID
		// (GET /users/{userId})
		GetUsersUserId(ctx echo.Context, userId int) error
		// Update User Information
		// (PATCH /users/{userId})
		PatchUsersUserId(ctx echo.Context, userId int) error
	}

	UserImpl struct {
		infra  infra.Service
		ucUser usecase.User
	}
)

func NewUser(infra infra.Service, ucUser usecase.User) User {
	return &UserImpl{
		infra:  infra,
		ucUser: ucUser,
	}
}

func (impl *UserImpl) PostUser(ctx echo.Context) error {
	var requestUser openapi.User
	if err := ctx.Bind(&requestUser); err != nil {
		return err
	}
	user, err := impl.ucUser.Create(ctx.Request().Context(), &ent.User{
		Name:  requestUser.Name,
		Email: cEnt.NullStringFromOapiEmail(requestUser.Email),
	})
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, response.ToUser(user))
}

func (impl *UserImpl) GetUsersUserId(ctx echo.Context, userId int) error {
	user, err := impl.ucUser.GetByUserID(ctx.Request().Context(), schema.UserID(userId))
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, response.ToUser(user))
}

func (impl *UserImpl) PatchUsersUserId(ctx echo.Context, userId int) error {
	params := make(map[string]interface{})
	if err := ctx.Bind(&params); err != nil {
		return err
	}
	user, err := impl.ucUser.Update(ctx.Request().Context(), schema.UserID(userId), params)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, response.ToUser(user))
}
