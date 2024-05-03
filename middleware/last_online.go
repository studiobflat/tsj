package middleware

import (
	"context"
	"net/http"
	"reflect"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type LastOnlineProvider interface {
	UpdateLastOnline(ctx context.Context, uid string) error
}

func UpdateLastOnline(skipper middleware.Skipper, lo LastOnlineProvider) echo.MiddlewareFunc {
	fbauth := func(next echo.HandlerFunc) echo.HandlerFunc {
		handler := func(c echo.Context) error {
			if skipper(c) {
				return next(c)
			}

			req := c.Request()
			uid := c.Get(UserIDContextKey)

			if uid == nil || reflect.TypeOf(uid).Name() != "string" {
				return echo.NewHTTPError(http.StatusBadRequest, `user id not found`)
			}

			_uid := uid.(string)
			if err := lo.UpdateLastOnline(req.Context(), _uid); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}

			return next(c)
		}

		return handler
	}
	return fbauth
}
