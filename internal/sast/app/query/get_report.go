package query

import (
	"context"

	"github.com/nikodem-kirsz/sast-rest-service/internal/common/decorator"
	"github.com/sirupsen/logrus"
)

type SpecifiedReport struct {
	UUID string
}

type GetReportHandler decorator.QueryHandler[SpecifiedReport, Report]

type getReportHandler struct {
	readModel GetReportReadModel
}

func NewGetReportHandler(
	readModel GetReportReadModel,
	logger *logrus.Entry,
) GetReportHandler {
	if readModel == nil {
		panic("nil readModel")
	}

	return decorator.ApplyQueryDecorators[SpecifiedReport, Report](
		getReportHandler{readModel: readModel},
		logger,
	)
}

type GetReportReadModel interface {
	GetReport(ctx context.Context, reportUUID string) (Report, error)
}

func (h getReportHandler) Handle(ctx context.Context, query SpecifiedReport) (re Report, err error) {
	return h.readModel.GetReport(ctx, query.UUID)
}
