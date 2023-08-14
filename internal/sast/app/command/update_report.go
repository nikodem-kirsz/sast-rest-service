package command

import (
	"context"
	"fmt"
	"time"

	"github.com/nikodem-kirsz/sast-rest-service/internal/common/decorator"
	"github.com/nikodem-kirsz/sast-rest-service/internal/common/logs"
	"github.com/nikodem-kirsz/sast-rest-service/internal/sast/domain/report"
	"github.com/sirupsen/logrus"
)

type UpdateReport struct {
	UUID          string
	Name          string
	Description   string
	Time          time.Time
	ReportContent string
}
type UpdateReportHandler decorator.CommandHandler[UpdateReport]

type updateReportHandler struct {
	repo report.Repository
}

func NewUpdateReportHandler(
	repo report.Repository,
	logger *logrus.Entry,
) UpdateReportHandler {
	if repo == nil {
		panic("nil repo")
	}

	return decorator.ApplyCommandDecorator[UpdateReport](
		updateReportHandler{repo: repo},
		logger,
	)
}

func (h updateReportHandler) Handle(ctx context.Context, cmd UpdateReport) (err error) {
	fmt.Printf("PUPAAAAAAAAAAAAAAAAAAAAA", cmd)
	defer func() {
		logs.LogCommandExecution("UpdateReport", cmd, err)
	}()

	fmt.Printf("Inside updates handle", cmd)

	re, err := report.NewReport(cmd.UUID, cmd.Name, cmd.Description, cmd.Time, cmd.ReportContent)
	if err := h.repo.UpdateReport(ctx, re); err != nil {
		return err
	}

	return nil
}
