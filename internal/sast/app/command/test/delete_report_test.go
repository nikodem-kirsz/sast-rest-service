package command

import (
	"context"
	"errors"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	command "github.com/nikodem-kirsz/sast-rest-service/internal/sast/app/command"
)

func TestDeleteReportHandler_Handle(t *testing.T) {
	logger := logrus.NewEntry(logrus.New())
	reportRepo := &MockReportRepository{}
	handler := command.NewDeleteReportHandler(reportRepo, logger)

	t.Run("Valid report deletion", func(t *testing.T) {
		ctx := context.Background()
		uuid := "some-uuid"

		reportRepo.DeleteReportFunc = func(ctx context.Context, id string) error {
			// Simulate successful report deletion
			return nil
		}

		err := handler.Handle(ctx, command.DeleteReport{UUID: uuid})
		assert.NoError(t, err)
	})

	t.Run("Report deletion error", func(t *testing.T) {
		ctx := context.Background()
		uuid := "some-uuid"

		expectedErr := errors.New("failed to delete report")
		reportRepo.DeleteReportFunc = func(ctx context.Context, id string) error {
			// Simulate report deletion error
			return expectedErr
		}

		err := handler.Handle(ctx, command.DeleteReport{UUID: uuid})
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
	})
}
