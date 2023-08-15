package command

import (
	"context"
	"errors"
	"testing"
	"time"

	command "github.com/nikodem-kirsz/sast-rest-service/internal/sast/app/command"
	"github.com/nikodem-kirsz/sast-rest-service/internal/sast/domain/report"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestCreateReportHandler_Handle(t *testing.T) {
	logger := logrus.NewEntry(logrus.New())
	reportRepo := &MockReportRepository{}
	handler := command.NewCreateReportHandler(reportRepo, logger)

	t.Run("Valid report creation", func(t *testing.T) {
		ctx := context.Background()
		cmd := command.CreateReport{
			UUID:          "some-uuid",
			Name:          "Test Report",
			Description:   "A test report",
			Time:          time.Now(),
			ReportContent: "Report content",
		}

		reportRepo.CreateReportFunc = func(ctx context.Context, r report.Report) error {
			// Simulate successful report creation
			return nil
		}

		err := handler.Handle(ctx, cmd)
		assert.NoError(t, err)
	})

	t.Run("Report creation error", func(t *testing.T) {
		ctx := context.Background()
		cmd := command.CreateReport{
			UUID:          "some-uuid",
			Name:          "Test Report",
			Description:   "A test report",
			Time:          time.Now(),
			ReportContent: "Report content",
		}

		expectedErr := errors.New("failed to create report")
		reportRepo.CreateReportFunc = func(ctx context.Context, r report.Report) error {
			// Simulate report creation error
			return expectedErr
		}

		err := handler.Handle(ctx, cmd)
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
	})
}
