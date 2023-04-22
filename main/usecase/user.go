package usecase

import (
	"context"
	"github.com/kondroid00/sample-server-2022/main/domain/ent"
	"github.com/kondroid00/sample-server-2022/main/domain/ent/schema"
	"github.com/kondroid00/sample-server-2022/main/domain/service"
	"github.com/kondroid00/sample-server-2022/main/infra"
)

type (
	User interface {
		GetByUserID(ctx context.Context, userID schema.UserID) (*ent.User, error)
		Create(ctx context.Context, user *ent.User) (*ent.User, error)
		Update(ctx context.Context, userID schema.UserID, params map[string]interface{}) (*ent.User, error)
	}

	UserImpl struct {
		infra  infra.Service
		svUser service.User
	}
)

func NewUser(infra infra.Service, svUser service.User) User {
	return &UserImpl{
		infra:  infra,
		svUser: svUser,
	}
}

func (impl *UserImpl) GetByUserID(ctx context.Context, userID schema.UserID) (*ent.User, error) {
	db := impl.infra.GetReadDB()
	return impl.svUser.GetByUserID(ctx, db.User, userID)
}

func (impl *UserImpl) Create(ctx context.Context, user *ent.User) (*ent.User, error) {
	db := impl.infra.GetDB()
	return impl.svUser.Create(ctx, db.User, user)
}

func (impl *UserImpl) Update(ctx context.Context, userID schema.UserID, params map[string]interface{}) (_ *ent.User, returnErr error) {
	tx, err := impl.infra.GetDB().Tx(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := tx.RollbackUnlessCommitted(); err != nil {
			returnErr = err
		}
	}()

	user, err := impl.svUser.Update(ctx, tx.User, userID, params)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return user, nil
}
