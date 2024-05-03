package vndcontext

import (
	"context"
	"fmt"
	"reflect"

	"github.com/labstack/echo/v4"
	tsjmiddleware "github.com/studiobflat/tsj/middleware"
)

type VndContext interface {
	echo.Context
	RequestContext() context.Context
	RequestId() string
	UserId() (string, error)
}

type VContext struct {
	echo.Context
}

func (c *VContext) RequestContext() context.Context {
	return c.Request().Context()
}

func (c *VContext) RequestId() string {
	id := c.Get(tsjmiddleware.RequestIDContextKey)
	if id != nil && reflect.TypeOf(id).Name() == "string" {
		return id.(string)
	}

	xid := c.Request().Header.Get(echo.HeaderXRequestID)
	if len(xid) > 0 {
		return xid
	}

	return ""
}

func (c *VContext) UserId() (string, error) {
	id := c.Get(tsjmiddleware.UserIDContextKey)
	if id != nil && reflect.TypeOf(id).Name() == "string" {
		return id.(string), nil
	}

	return "", fmt.Errorf(`user id not found`)
}
