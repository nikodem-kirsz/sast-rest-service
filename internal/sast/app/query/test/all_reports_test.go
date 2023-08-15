package query

import (
	"context"
	"errors"
	"testing"

	query "github.com/nikodem-kirsz/sast-rest-service/internal/sast/app/query"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

type MockAllReportsReadModel struct {
	GetAllReportsFunc func(ctx context.Context) ([]query.Report, error)
}

func (m *MockAllReportsReadModel) GetAllReports(ctx context.Context) ([]query.Report, error) {
	if m.GetAllReportsFunc != nil {
		return m.GetAllReportsFunc(ctx)
	}
	return nil, nil
}

func TestAllReportsHandler_Handle(t *testing.T) {
	logger := logrus.NewEntry(logrus.New())
	readModel := &MockAllReportsReadModel{}
	handler := query.NewAllReportsHandler(readModel, logger)

	t.Run("Valid report retrieval", func(t *testing.T) {
		ctx := context.Background()
		expectedReports := []query.Report{
			{UUID: "1", Name: "Report 1"},
			{UUID: "2", Name: "Report 2"},
		}

		readModel.GetAllReportsFunc = func(ctx context.Context) ([]query.Report, error) {
			// Simulate successful report retrieval
			return expectedReports, nil
		}

		reports, err := handler.Handle(ctx, query.AllReports{})
		assert.NoError(t, err)
		assert.Equal(t, expectedReports, reports)
	})

	t.Run("Report retrieval error", func(t *testing.T) {
		ctx := context.Background()

		expectedErr := errors.New("failed to retrieve reports")
		readModel.GetAllReportsFunc = func(ctx context.Context) ([]query.Report, error) {
			// Simulate report retrieval error
			return nil, expectedErr
		}

		reports, err := handler.Handle(ctx, query.AllReports{})
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		assert.Nil(t, reports)
	})
}
