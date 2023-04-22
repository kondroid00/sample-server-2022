package middleware

import (
	"github.com/kondroid00/sample-server-2022/package/errorcode"
	"github.com/kondroid00/sample-server-2022/package/errors"
	"github.com/labstack/echo/v4"
)

func NewHttpStatusInterceptor() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err == nil {
				return nil
			}

			var httpError *errorcode.HttpError
			if errors.As(err, &httpError) {
				return c.JSON(httpError.HttpStatusCode(), httpError.ErrorCode())
			}
			return err
		}
	}
}
