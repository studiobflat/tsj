package route

import (
	"example/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/studiobflat/tsj/rest"
	"github.com/studiobflat/tsj/runner"
)

type V1 struct {
	*echo.Group
}

func (v1 *V1) Configure(rn *runner.Runner) error {
	s := service.NewService()
	return v1.registerRoutes(s)
}

func (v1 *V1) registerRoutes(s *service.Service) error {
	v1.GET("/", rest.Wrapper(s.GetSuccess))

	return nil
}
