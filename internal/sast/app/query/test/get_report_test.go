package query

import (
	"context"
	"errors"
	"testing"

	query "github.com/nikodem-kirsz/sast-rest-service/internal/sast/app/query"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

type MockGetReportReadModel struct {
	GetReportFunc func(ctx context.Context, reportUUID string) (query.Report, error)
}

func (m *MockGetReportReadModel) GetReport(ctx context.Context, reportUUID string) (query.Report, error) {
	if m.GetReportFunc != nil {
		return m.GetReportFunc(ctx, reportUUID)
	}
	return query.Report{}, nil
}

func TestGetReportHandler_Handle(t *testing.T) {
	logger := logrus.NewEntry(logrus.New())
	readModel := &MockGetReportReadModel{}
	handler := query.NewGetReportHandler(readModel, logger)

	t.Run("Valid report retrieval", func(t *testing.T) {
		ctx := context.Background()
		uuid := "some-uuid"
		expectedReport := query.Report{
			UUID: "some-uuid",
			Name: "Test Report",
		}

		readModel.GetReportFunc = func(ctx context.Context, reportUUID string) (query.Report, error) {
			// Simulate successful report retrieval
			return expectedReport, nil
		}

		report, err := handler.Handle(ctx, query.SpecifiedReport{UUID: uuid})
		assert.NoError(t, err)
		assert.Equal(t, expectedReport, report)
	})

	t.Run("Report retrieval error", func(t *testing.T) {
		ctx := context.Background()
		uuid := "some-uuid"

		expectedErr := errors.New("failed to retrieve report")
		readModel.GetReportFunc = func(ctx context.Context, reportUUID string) (query.Report, error) {
			// Simulate report retrieval error
			return query.Report{}, expectedErr
		}

		report, err := handler.Handle(ctx, query.SpecifiedReport{UUID: uuid})
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		assert.Equal(t, query.Report{}, report)
	})
}
