package service

import (
	"context"
	"github.com/kondroid00/sample-server-2022/main/domain/ent"
	"github.com/kondroid00/sample-server-2022/main/domain/ent/schema"
	"github.com/kondroid00/sample-server-2022/main/domain/repository"
	"github.com/kondroid00/sample-server-2022/main/infra"
)

type (
	User interface {
		GetByUserID(ctx context.Context, client *ent.UserClient, userID schema.UserID) (*ent.User, error)
		Create(ctx context.Context, client *ent.UserClient, user *ent.User) (*ent.User, error)
		Update(ctx context.Context, client *ent.UserClient, userID schema.UserID, params map[string]interface{}) (*ent.User, error)
	}

	UserImpl struct {
		infra    infra.Service
		repoUser repository.User
	}
)

func NewUser(infra infra.Service, repoUser repository.User) User {
	return &UserImpl{
		infra:    infra,
		repoUser: repoUser,
	}
}

func (impl *UserImpl) GetByUserID(ctx context.Context, client *ent.UserClient, userID schema.UserID) (*ent.User, error) {
	return impl.repoUser.Get(ctx, client, userID)
}

func (impl *UserImpl) Create(ctx context.Context, client *ent.UserClient, user *ent.User) (*ent.User, error) {
	user.State = schema.Enable
	return impl.repoUser.Create(ctx, client, user)
}

func (impl *UserImpl) Update(ctx context.Context, client *ent.UserClient, userID schema.UserID, params map[string]interface{}) (_ *ent.User, returnErr error) {
	return impl.repoUser.Update(ctx, client, userID, params)
}
