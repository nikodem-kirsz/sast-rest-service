package ports

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/nikodem-kirsz/sast-rest-service/internal/common/server/httperr"
	"github.com/nikodem-kirsz/sast-rest-service/internal/sast/app"
	"github.com/nikodem-kirsz/sast-rest-service/internal/sast/app/command"
	"github.com/nikodem-kirsz/sast-rest-service/internal/sast/app/query"
)

type HttpServer struct {
	app app.Application
}

func NewHttpServer(app app.Application) HttpServer {
	return HttpServer{app}
}

func (h HttpServer) GetSastReports(w http.ResponseWriter, r *http.Request) {
	var appReports []query.Report
	appReports, err := h.app.Queries.AllReports.Handle(r.Context(), query.AllReports{})
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}
	reports := appReportsToResponse(appReports)
	reportsResp := SastReports{reports}

	fmt.Println("Inside GetSastReports")
	fmt.Println("Reportsresp", reportsResp)
	render.Respond(w, r, reportsResp)
}

func (h HttpServer) CreateSastReport(w http.ResponseWriter, r *http.Request) {
	createReport := CreateSastReport{}
	if err := render.Decode(r, &createReport); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}

	cmd := command.CreateReport{
		UUID:          uuid.New().String(),
		Name:          createReport.Name,
		Description:   createReport.Description,
		Time:          createReport.Time,
		ReportContent: createReport.ReportContent,
	}

	err := h.app.Commands.CreateReport.Handle(r.Context(), cmd)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	w.Header().Set("content-location", "/sast-reports/"+cmd.UUID)
	fmt.Println("Inside CreateSastReports")
	w.WriteHeader(http.StatusNoContent)
}

func (h HttpServer) DeleteReport(w http.ResponseWriter, r *http.Request, reportUUID uuid.UUID) {
	err := h.app.Commands.DeleteReport.Handle(r.Context(), command.DeleteReport{
		UUID: reportUUID.String(),
	})
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}
	w.Header().Set("content-location", "/sast-reports/"+reportUUID.String())
	fmt.Println("Inside DeletSastReports")
	w.WriteHeader(http.StatusNoContent)
}

func (h HttpServer) GetReport(w http.ResponseWriter, r *http.Request, reportUUID uuid.UUID) {
	var appReport query.Report
	appReport, err := h.app.Queries.GetReport.Handle(r.Context(), query.SpecifiedReport{
		UUID: reportUUID.String(),
	})
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}
	report := appReportToResponse(appReport)
	reportResp := SastReport{
		Uuid:          report.Uuid,
		Name:          report.Name,
		Description:   report.Description,
		Time:          report.Time,
		ReportContent: report.ReportContent,
	}
	w.Header().Set("content-location", "/sast-reports/"+reportUUID.String())
	fmt.Println("Inside GetReport")
	render.Respond(w, r, reportResp)
}

func (h HttpServer) UpdateReport(w http.ResponseWriter, r *http.Request, reportUUID uuid.UUID) {
	updateReport := UpdateSastReport{}
	if err := render.Decode(r, &updateReport); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}

	cmd := command.UpdateReport{
		UUID:          reportUUID.String(),
		Name:          *updateReport.Name,
		Description:   *updateReport.Description,
		Time:          *updateReport.Time,
		ReportContent: *updateReport.ReportContent,
	}
	err := h.app.Commands.UpdateReport.Handle(r.Context(), cmd)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	w.Header().Set("content-location", "/sast-reports/"+reportUUID.String())
	fmt.Println("Inside UpdateSastReports")
	w.WriteHeader(http.StatusNoContent)
}

func appReportsToResponse(appReports []query.Report) []SastReport {
	var reports []SastReport
	for _, tm := range appReports {
		r := appReportToResponse(tm)

		reports = append(reports, r)
	}

	return reports
}

func appReportToResponse(appReport query.Report) SastReport {
	return SastReport{
		Uuid:          uuid.MustParse(appReport.UUID),
		Name:          appReport.Name,
		Description:   appReport.Description,
		Time:          appReport.Time,
		ReportContent: appReport.ReportContent,
	}
}
