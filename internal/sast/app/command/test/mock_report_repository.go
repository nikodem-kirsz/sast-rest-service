package command

import (
	"context"

	"github.com/nikodem-kirsz/sast-rest-service/internal/sast/domain/report"
)

// MockReportRepository is a mock implementation of the report.Repository interface
type MockReportRepository struct {
	CreateReportFunc func(ctx context.Context, r report.Report) error
	DeleteReportFunc func(ctx context.Context, reportUUID string) error
	UpdateReportFunc func(ctx context.Context, re report.Report) error
}

// DeleteReport implements report.Repository.
func (m *MockReportRepository) DeleteReport(ctx context.Context, reportUUID string) error {
	if m.DeleteReportFunc != nil {
		return m.DeleteReportFunc(ctx, reportUUID)
	}
	return nil
}

// UpdateReport implements report.Repository.
func (m *MockReportRepository) UpdateReport(ctx context.Context, r *report.Report) error {
	if m.UpdateReportFunc != nil {
		return m.UpdateReportFunc(ctx, *r)
	}
	return nil
}

func (m *MockReportRepository) CreateReport(ctx context.Context, r *report.Report) error {
	if m.CreateReportFunc != nil {
		return m.CreateReportFunc(ctx, *r)
	}
	return nil
}
