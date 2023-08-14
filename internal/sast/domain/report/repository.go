package report

import (
	"context"
	"fmt"
)

type NotFoundError struct {
	ReportUUID string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("report '%s' not found", e.ReportUUID)
}

type Repository interface {
	CreateReport(ctx context.Context, re *Report) error
	DeleteReport(ctx context.Context, reportUUID string) error
	UpdateReport(ctx context.Context, re *Report) error
}
