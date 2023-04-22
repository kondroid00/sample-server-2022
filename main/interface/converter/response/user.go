package response

import (
	"github.com/kondroid00/sample-server-2022/main/domain/ent"
	cEnt "github.com/kondroid00/sample-server-2022/main/interface/converter/ent"
	"github.com/kondroid00/sample-server-2022/main/interface/openapi"
)

func ToUser(user *ent.User) *openapi.User {
	return &openapi.User{
		Id:    cEnt.Int64ToPtr(int64(user.ID)),
		Name:  user.Name,
		Email: cEnt.NullStringToEmailPtr(user.Email),
		State: cEnt.StringToPtr(string(user.State)),
	}
}
