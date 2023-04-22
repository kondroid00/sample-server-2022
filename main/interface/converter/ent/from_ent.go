package ent

import (
	"entgo.io/ent/dialect/sql"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

func Int64ToPtr(i int64) *int64 {
	if i == 0 {
		return nil
	}
	return &i
}

func StringToPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func NullStringToPtr(s sql.NullString) *string {
	if s.Valid == false {
		return nil
	}
	return &(s.String)
}

func NullStringToEmailPtr(s sql.NullString) *openapi_types.Email {
	if s.Valid == false {
		return nil
	}
	if s.String == "" {
		return nil
	}
	e := openapi_types.Email(s.String)
	return &e
}
