package middleware

import (
	"github.com/labstack/echo/v4"
	vnderror "github.com/studiobflat/tsj/error"
)

func ErrorHandler(next echo.HTTPErrorHandler) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		// same check as default handler
		if c.Response().Committed {
			return
		}

		if e, ok := err.(*vnderror.Error); ok {
			c.JSON(e.Code, e)
			return
		}

		if next != nil {
			next(err, c)
		}
	}
}
