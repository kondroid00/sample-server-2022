package ent

import (
	"entgo.io/ent/dialect/sql"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

func NullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{}
	}
	return sql.NullString{
		String: *s,
		Valid:  true,
	}
}

func NullInt64(i *int64) sql.NullInt64 {
	if i == nil {
		return sql.NullInt64{}
	}
	return sql.NullInt64{
		Int64: *i,
		Valid: true,
	}
}

func NullStringFromOapiEmail(e *openapi_types.Email) sql.NullString {
	if e == nil {
		return sql.NullString{}
	}
	return sql.NullString{
		String: string(*e),
		Valid:  true,
	}
}
