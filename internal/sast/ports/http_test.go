package ports_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/nikodem-kirsz/sast-rest-service/internal/sast/app"
	"github.com/nikodem-kirsz/sast-rest-service/internal/sast/app/command"
	"github.com/nikodem-kirsz/sast-rest-service/internal/sast/app/query"
	"github.com/nikodem-kirsz/sast-rest-service/internal/sast/ports"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ApplicationMock struct {
	mock.Mock
}

func (m *ApplicationMock) ExecuteCommand(ctx context.Context, cmd app.Command) error {
	args := m.Called(ctx, cmd)
	return args.Error(0)
}

func (m *ApplicationMock) ExecuteQuery(ctx context.Context, q app.Query) (interface{}, error) {
	args := m.Called(ctx, q)
	return args.Get(0), args.Error(1)
}

func (m *ApplicationMock) GetCommands() app.CommandHandlers {
	return app.CommandHandlers{
		command.CommandTypeCreateReport: &MockCreateReportCommand{},
		// Add more command mock objects here as needed
	}
}

func (m *ApplicationMock) GetQueries() app.QueryHandlers {
	return app.QueryHandlers{
		query.QueryTypeAllReports: &MockAllReportsQuery{},
		// Add more query mock objects here as needed
	}
}

type MockCreateReportCommand struct {
	mock.Mock
}

func (m *MockCreateReportCommand) Handle(ctx context.Context, cmd command.CreateReport) error {
	args := m.Called(ctx, cmd)
	return args.Error(0)
}

type MockAllReportsQuery struct {
	mock.Mock
}

func (m *MockAllReportsQuery) Handle(ctx context.Context, q query.AllReports) ([]query.Report, error) {
	args := m.Called(ctx, q)
	return args.Get(0).([]query.Report), args.Error(1)
}

func TestHttpServer_CreateSastReport(t *testing.T) {
	appMock := &app.ApplicationMock{}
	server := ports.NewHttpServer(appMock)

	createReport := ports.CreateSastReport{
		Name:          "Test Report",
		Description:   "Test Description",
		Time:          time.Now(),
		ReportContent: "Test Content",
	}
	requestBody, _ := json.Marshal(createReport)
	request := httptest.NewRequest(http.MethodPost, "/sast-reports", bytes.NewReader(requestBody))
	recorder := httptest.NewRecorder()

	appMock.Commands.CreateReport.On("Handle", request.Context(), command.CreateReport{
		Name:          createReport.Name,
		Description:   createReport.Description,
		Time:          createReport.Time,
		ReportContent: createReport.ReportContent,
	}).Return(nil)

	server.CreateSastReport(recorder, request)

	assert.Equal(t, http.StatusCreated, recorder.Code)
	// Add more assertions if needed
}

func TestHttpServer_GetSastReports(t *testing.T) {
	appMock := &app.ApplicationMock{}
	server := ports.NewHttpServer(appMock)

	request := httptest.NewRequest(http.MethodGet, "/sast-reports", nil)
	recorder := httptest.NewRecorder()

	reports := []query.Report{
		{
			UUID:          uuid.New().String(),
			Name:          "Test Report 1",
			Description:   "Description 1",
			Time:          time.Now(),
			ReportContent: "Content 1",
		},
		// Add more test reports as needed
	}
	appMock.Queries.AllReports.On("Handle", request.Context(), query.AllReports{}).Return(reports, nil)

	server.GetSastReports(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	// Add more assertions if needed
}

// Write similar tests for other HttpServer methods: DeleteReport, GetReport, UpdateReport
