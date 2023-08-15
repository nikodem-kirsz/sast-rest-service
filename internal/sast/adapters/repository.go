package adapters

import (
	"context"

	"github.com/nikodem-kirsz/sast-rest-service/internal/sast/app/query"
	"github.com/nikodem-kirsz/sast-rest-service/internal/sast/domain/report"
)

type Repository interface {
	CreateReport(ctx context.Context, re *report.Report) error
	GetReport(ctx context.Context, reportUUID string) (query.Report, error)
	DeleteReport(ctx context.Context, reportUUID string) error
	UpdateReport(ctx context.Context, updatedReport *report.Report) error
	GetAllReports(ctx context.Context) ([]query.Report, error)
}
