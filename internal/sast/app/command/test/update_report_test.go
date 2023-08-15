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

func TestUpdateReportHandler_Handle(t *testing.T) {
	logger := logrus.NewEntry(logrus.New())
	reportRepo := &MockReportRepository{}
	handler := command.NewUpdateReportHandler(reportRepo, logger)

	t.Run("Valid report update", func(t *testing.T) {
		ctx := context.Background()
		cmd := command.UpdateReport{
			UUID:          "some-uuid",
			Name:          "Updated Report",
			Description:   "An updated report",
			Time:          time.Now(),
			ReportContent: "Updated report content",
		}

		reportRepo.UpdateReportFunc = func(ctx context.Context, r report.Report) error {
			// Simulate successful report update
			return nil
		}

		err := handler.Handle(ctx, cmd)
		assert.NoError(t, err)
	})

	t.Run("Report update error", func(t *testing.T) {
		ctx := context.Background()
		cmd := command.UpdateReport{
			UUID:          "some-uuid",
			Name:          "Updated Report",
			Description:   "An updated report",
			Time:          time.Now(),
			ReportContent: "Updated report content",
		}

		expectedErr := errors.New("failed to update report")
		reportRepo.UpdateReportFunc = func(ctx context.Context, r report.Report) error {
			// Simulate report update error
			return expectedErr
		}

		err := handler.Handle(ctx, cmd)
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
	})
}
