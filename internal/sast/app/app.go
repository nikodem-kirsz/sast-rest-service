package app

import (
	"context"

	"github.com/nikodem-kirsz/sast-rest-service/internal/sast/app/query"
	"github.com/nikodem-kirsz/sast-rest-service/internal/sast/domain/report"
)

type Repository interface {
	AddReport(ctx context.Context, re *report.Report) error
	GetReport(ctx context.Context, reportUUID string) (*report.Report, error)
	DeleteReport(ctx context.Context, reportUUID string) error
	UpdateReport(
		ctx context.Context,
		reportUUID string,
		updateFn func(ctx context.Context, re *report.Report) (*report.Report, error),
	) error
	GetAllReports(ctx context.Context) ([]query.Report, error)
}
type Application struct {
	// repository *Repository
	Queries Queries
}

// func NewApplication(repository Repository) *Application {
// 	return &Application{
// 		repository: &repository,
// 	}
// }

type Queries struct {
	AllReports query.AllReportsHandler
}
