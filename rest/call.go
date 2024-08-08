package rest

import (
	"github.com/studiobflat/tsj/logger"
	"github.com/studiobflat/tsj/vndcontext"
)

type CallDelegate[REQ any] func(*logger.Logger, vndcontext.VndContext, *REQ) (*Result, error)

func Call[REQ any](e vndcontext.VndContext, req *REQ, name string, delegate CallDelegate[REQ]) (*Result, error) {
	log := logger.GetLogger(name)
	defer func() {
		log.Infow("completed")
		log.Sync()
	}()

	requestId := e.RequestId()
	log.With([]interface{}{
		"request_id", requestId,
	}...)

	return delegate(log, e, req)
}

type CallDelegateAny[REQ any] func(*logger.Logger, vndcontext.VndContext, *REQ) (any, error)

func CallAny[REQ any](e vndcontext.VndContext, req *REQ, name string, delegate CallDelegateAny[REQ]) (any, error) {
	log := logger.GetLogger(name)
	defer func() {
		log.Infow("completed")
		log.Sync()
	}()

	requestId := e.RequestId()
	log.With([]interface{}{
		"request_id", requestId,
	}...)

	return delegate(log, e, req)
}
