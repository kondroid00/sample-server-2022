package repository

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/kondroid00/sample-server-2022/main/domain/ent"
	"github.com/kondroid00/sample-server-2022/main/domain/ent/schema"
	"github.com/kondroid00/sample-server-2022/main/domain/ent/user"
	"github.com/kondroid00/sample-server-2022/main/server/errorcode"
	pkgErrCode "github.com/kondroid00/sample-server-2022/package/errorcode"
	"github.com/kondroid00/sample-server-2022/package/errors"
	"net/http"
)

type (
	User interface {
		Get(ctx context.Context, client *ent.UserClient, userID schema.UserID) (*ent.User, error)
		Create(ctx context.Context, client *ent.UserClient, user *ent.User) (*ent.User, error)
		Update(ctx context.Context, client *ent.UserClient, userID schema.UserID, params map[string]interface{}) (*ent.User, error)
	}

	UserImpl struct{}
)

func NewUser() User {
	return &UserImpl{}
}

func (impl *UserImpl) Get(ctx context.Context, client *ent.UserClient, userID schema.UserID) (*ent.User, error) {
	returnUser, err := client.Get(ctx, userID)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return nil, pkgErrCode.NewHttpError(
				http.StatusNotFound,
				errorcode.NOT_FOUND.SetParam(fmt.Sprint(userID)),
				errors.Stack(err),
			)
		default:
			return nil, errors.Stack(err)
		}
	}
	return returnUser, nil
}

func (impl *UserImpl) Create(ctx context.Context, client *ent.UserClient, user *ent.User) (*ent.User, error) {
	returnUser, err := client.
		Create().
		SetName(user.Name).
		SetEmail(user.Email).
		SetState(user.State).
		Save(ctx)
	if err != nil {
		return nil, errors.Stack(err)
	}
	return returnUser, nil
}

func (impl *UserImpl) Update(ctx context.Context, client *ent.UserClient, userID schema.UserID, params map[string]interface{}) (*ent.User, error) {
	u := client.UpdateOneID(userID)
	mError := errors.NewMultiError()
	for k, v := range params {
		switch k {
		case user.FieldName:
			mError.Append(setValue(v, func(value string) error {
				if value == "" {
					return pkgErrCode.NewHttpError(
						http.StatusBadRequest,
						errorcode.INVALID_PARAMS.SetParam(fmt.Sprint(k)),
						errors.New("invalid params"),
					)
				}
				u.SetName(value)
				return nil
			}))
		case user.FieldEmail:
			mError.Append(setNull(v, func(value string) error {
				if value == "" {
					u.SetEmail(sql.NullString{})
				} else {
					u.SetEmail(sql.NullString{String: value, Valid: true})
				}
				return nil
			}))
		case user.FieldState:
			mError.Append(setNull(v, func(value schema.UserState) error {
				u.SetState(value)
				return nil
			}))
		}
	}
	if err := mError.GetError(); err != nil {
		return nil, errors.Stack(err)
	}

	returnUser, err := u.Save(ctx)
	if err != nil {
		return nil, errors.Stack(err)
	}
	return returnUser, nil
}
