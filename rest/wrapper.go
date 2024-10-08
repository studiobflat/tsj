package rest

import (
	"net/http"
	"reflect"
	"runtime"
	"time"

	"github.com/labstack/echo/v4"
	vnderror "github.com/studiobflat/tsj/error"
	"github.com/studiobflat/tsj/logger"
	"github.com/studiobflat/tsj/vndcontext"
)

const RequestObjectContextKey = "service_requestObject"

func Wrapper[TREQ any](wrapped func(vndcontext.VndContext, *TREQ) (*Result, error)) echo.HandlerFunc {
	return func(c echo.Context) error {
		log := logger.GetLogger("Wrapper")
		defer log.Sync()

		vndc := c.(*vndcontext.VContext)
		requestId := vndc.RequestId()
		handler := runtime.FuncForPC(reflect.ValueOf(wrapped).Pointer()).Name()
		log.Infow("request begin",
			"request_id", requestId,
			"at", time.Now().Format(time.RFC3339),
			"path", c.Request().RequestURI,
			"handler", handler,
		)

		var req TREQ
		if err := c.Bind(&req); err != nil {
			log.Errorw("fail to bind request", "request_uri", c.Request().RequestURI, "err", err)
			return &vnderror.Error{
				CustomCode: -40001,
				HTTPError: &echo.HTTPError{
					Code:    http.StatusBadRequest,
					Message: "invalid request",
				},
			}
		}

		if err := c.Validate(&req); err != nil {
			log.Errorw("fail to validate request", "request_uri", c.Request().RequestURI, "request_object", req, "err", err)
			return &vnderror.Error{
				CustomCode: -40002,
				HTTPError: &echo.HTTPError{
					Code:    http.StatusBadRequest,
					Message: "invalid request",
				},
			}
		}

		c.Set(RequestObjectContextKey, req)

		res, err := wrapped(vndc, &req)
		if err != nil {
			log.Errorw("request end with error", "request_id", requestId, "at", time.Now().Format(time.RFC3339), "err", err)
			return err
		}

		status := c.Response().Status
		if status != 0 {
			log.Infow("request end", "request_id", requestId, "at", time.Now().Format(time.RFC3339), "status", status)

			return c.JSON(status, res)
		}

		log.Infow("request end", "request_id", requestId, "at", time.Now().Format(time.RFC3339), "status", http.StatusOK)

		return c.JSON(http.StatusOK, res)
	}
}

func WrapperAny[TREQ any](wrapped func(vndcontext.VndContext, *TREQ) (any, error)) echo.HandlerFunc {
	return func(c echo.Context) error {
		log := logger.GetLogger("Wrapper")
		defer log.Sync()

		vndc := c.(*vndcontext.VContext)
		requestId := vndc.RequestId()
		handler := runtime.FuncForPC(reflect.ValueOf(wrapped).Pointer()).Name()
		log.Infow("request begin",
			"request_id", requestId,
			"at", time.Now().Format(time.RFC3339),
			"path", c.Request().RequestURI,
			"handler", handler,
		)

		var req TREQ
		if err := c.Bind(&req); err != nil {
			log.Errorw("fail to bind request", "request_uri", c.Request().RequestURI, "err", err)
			return &vnderror.Error{
				CustomCode: -40001,
				HTTPError: &echo.HTTPError{
					Code:    http.StatusBadRequest,
					Message: "invalid request",
				},
			}
		}

		if err := c.Validate(&req); err != nil {
			log.Errorw("fail to validate request", "request_uri", c.Request().RequestURI, "request_object", req, "err", err)
			return &vnderror.Error{
				CustomCode: -40002,
				HTTPError: &echo.HTTPError{
					Code:    http.StatusBadRequest,
					Message: "invalid request",
				},
			}
		}

		c.Set(RequestObjectContextKey, req)

		res, err := wrapped(vndc, &req)
		if err != nil {
			log.Errorw("request end with error", "request_id", requestId, "at", time.Now().Format(time.RFC3339), "err", err)
			return err
		}

		status := c.Response().Status
		if status != 0 {
			log.Infow("request end", "request_id", requestId, "at", time.Now().Format(time.RFC3339), "status", status)

			return c.JSON(status, res)
		}

		log.Infow("request end", "request_id", requestId, "at", time.Now().Format(time.RFC3339), "status", http.StatusOK)

		return c.JSON(http.StatusOK, res)
	}
}

func WrapperSSE[TREQ any](wrapped func(vndcontext.VndContext, *TREQ) error) echo.HandlerFunc {
	return func(c echo.Context) error {
		log := logger.GetLogger("Wrapper")
		defer log.Sync()

		vndc := c.(*vndcontext.VContext)
		requestId := vndc.RequestId()
		handler := runtime.FuncForPC(reflect.ValueOf(wrapped).Pointer()).Name()
		log.Infow("request begin",
			"request_id", requestId,
			"at", time.Now().Format(time.RFC3339),
			"path", c.Request().RequestURI,
			"handler", handler,
		)

		var req TREQ
		if err := c.Bind(&req); err != nil {
			log.Errorw("fail to bind request", "request_uri", c.Request().RequestURI, "err", err)
			return &vnderror.Error{
				CustomCode: -40001,
				HTTPError: &echo.HTTPError{
					Code:    http.StatusBadRequest,
					Message: "invalid request",
				},
			}
		}

		if err := c.Validate(&req); err != nil {
			log.Errorw("fail to validate request", "request_uri", c.Request().RequestURI, "request_object", req, "err", err)
			return &vnderror.Error{
				CustomCode: -40002,
				HTTPError: &echo.HTTPError{
					Code:    http.StatusBadRequest,
					Message: "invalid request",
				},
			}
		}

		c.Set(RequestObjectContextKey, req)

		err := wrapped(vndc, &req)
		if err != nil {
			log.Errorw("request end with error", "request_id", requestId, "at", time.Now().Format(time.RFC3339), "err", err)
			return err
		}

		log.Infow("request end", "request_id", requestId, "at", time.Now().Format(time.RFC3339), "status", http.StatusOK)

		return nil
	}
}
