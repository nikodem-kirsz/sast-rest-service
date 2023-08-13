package query

import (
	"context"

	"github.com/nikodem-kirsz/sast-rest-service/internal/common/decorator"
	"github.com/sirupsen/logrus"
)

type AllReports struct{}

type AllReportsHandler decorator.QueryHandler[AllReports, []Report]

type allReportsHandler struct {
	readModel AllReportsReadModel
}

func NewAllReportsHandler(
	readModel AllReportsReadModel,
	logger *logrus.Entry,
) AllReportsHandler {
	if readModel == nil {
		panic("nil readModel")
	}

	return decorator.ApplyQueryDecorators[AllReports, []Report](
		allReportsHandler{readModel: readModel},
		logger,
	)
}

type AllReportsReadModel interface {
	GetAllReports(ctx context.Context) ([]Report, error)
}

func (h allReportsHandler) Handle(ctx context.Context, _ AllReports) (re []Report, err error) {
	return h.readModel.GetAllReports(ctx)
}
