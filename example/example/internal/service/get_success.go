package service

import (
	_ "github.com/studiobflat/tsj/error"
	"github.com/studiobflat/tsj/logger"
	"github.com/studiobflat/tsj/rest"
	"github.com/studiobflat/tsj/vndcontext"
)

type GetSuccessReq struct {
}

func (s *Service) GetSuccess(e vndcontext.VndContext, req *GetSuccessReq) (*rest.Result, error) {
	delegate := func(log *logger.Logger, ctx vndcontext.VndContext, req *GetSuccessReq) (*rest.Result, error) {
		exec := NewGetSuccess(log)
		return exec.Execute(ctx, req)
	}
	return rest.Call[GetSuccessReq](e, req, "GetSuccess", delegate)
}

type getSuccess struct {
	log *logger.Logger
}

func NewGetSuccess(log *logger.Logger) *getSuccess {
	return &getSuccess{
		log: log,
	}
}

func (s *getSuccess) Execute(ctx vndcontext.VndContext, req *GetSuccessReq) (*rest.Result, error) {
	uid, err := ctx.UserId()
	if err != nil {
		s.log.Errorw("failed to get user id", "error", err)
		return nil, err
	}

	return &rest.Result{
		Data: uid,
	}, nil
}
