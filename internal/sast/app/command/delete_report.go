package command

import (
	"context"

	"github.com/nikodem-kirsz/sast-rest-service/internal/common/decorator"
	"github.com/nikodem-kirsz/sast-rest-service/internal/common/logs"
	"github.com/nikodem-kirsz/sast-rest-service/internal/sast/domain/report"
	"github.com/sirupsen/logrus"
)

type DeleteReport struct {
	UUID string
}

type DeleteReportHandler decorator.CommandHandler[DeleteReport]

type deleteReportHandler struct {
	repo report.Repository
}

func NewDeleteReportHandler(
	repo report.Repository,
	logger *logrus.Entry,
) DeleteReportHandler {
	if repo == nil {
		panic("nil repo")
	}

	return decorator.ApplyCommandDecorator[DeleteReport](
		deleteReportHandler{repo: repo},
		logger,
	)
}

func (h deleteReportHandler) Handle(ctx context.Context, cmd DeleteReport) (err error) {
	defer func() {
		logs.LogCommandExecution("DeleteReport", cmd, err)
	}()

	if err := h.repo.DeleteReport(ctx, cmd.UUID); err != nil {
		return err
	}

	return nil
}
