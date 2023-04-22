package repository

import (
	"fmt"
	"github.com/kondroid00/sample-server-2022/package/errors"
)

func setValue[T any](v interface{}, f func(T) error) error {
	if v == nil {
		return errors.New(fmt.Sprint("value is null"))
	}

	value, ok := v.(T)
	if !ok {
		return errors.New(fmt.Sprintf("cast error"))
	}

	return f(value)
}

func setNull[T any](v interface{}, f func(T) error) error {
	if v == nil {
		return nil
	}

	value, ok := v.(T)
	if !ok {
		return errors.New(fmt.Sprintf("cast error"))
	}

	return f(value)
}
