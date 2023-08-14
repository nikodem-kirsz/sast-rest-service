package command

import (
	"context"
	"time"

	"github.com/nikodem-kirsz/sast-rest-service/internal/common/decorator"
	"github.com/nikodem-kirsz/sast-rest-service/internal/common/logs"
	"github.com/nikodem-kirsz/sast-rest-service/internal/sast/domain/report"
	"github.com/sirupsen/logrus"
)

type CreateReport struct {
	UUID          string
	Name          string
	Description   string
	Time          time.Time
	ReportContent string
}

type CreateReportHandler decorator.CommandHandler[CreateReport]

type createReportHandler struct {
	repo report.Repository
}

func NewCreateReportHandler(
	repo report.Repository,
	logger *logrus.Entry,
) CreateReportHandler {
	if repo == nil {
		panic("nil repo")
	}

	return decorator.ApplyCommandDecorator[CreateReport](
		createReportHandler{repo: repo},
		logger,
	)
}

func (h createReportHandler) Handle(ctx context.Context, cmd CreateReport) (err error) {
	defer func() {
		logs.LogCommandExecution("CreateReport", cmd, err)
	}()

	re, err := report.NewReport(cmd.UUID, cmd.Name, cmd.Description, cmd.Time, cmd.ReportContent)
	if err != nil {
		return err
	}

	if err := h.repo.CreateReport(ctx, re); err != nil {
		return err
	}

	return nil
}
