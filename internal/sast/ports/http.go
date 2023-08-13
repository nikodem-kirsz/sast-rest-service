package ports

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/nikodem-kirsz/sast-rest-service/internal/common/server/httperr"
	"github.com/nikodem-kirsz/sast-rest-service/internal/sast/app"
	"github.com/nikodem-kirsz/sast-rest-service/internal/sast/app/query"
)

type HttpServer struct {
	app app.Application
}

func NewHttpServer(app app.Application) HttpServer {
	return HttpServer{app}
}

func (h HttpServer) GetSastReports(w http.ResponseWriter, r *http.Request) {
	// var reports []query.Report
	allReports, err := h.app.Queries.AllReports.Handle(r.Context(), query.AllReports{})
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}
	// reports = appReportsToResponse(allReports)

	fmt.Println("Inside GetSastReports")
	render.Respond(w, r, allReports)
}

func (h HttpServer) CreateSastReport(w http.ResponseWriter, r *http.Request) {

}

func (h HttpServer) DeleteReport(w http.ResponseWriter, r *http.Request, reportUUID uuid.UUID) {

}

func (h HttpServer) GetReport(w http.ResponseWriter, r *http.Request, reportUUID uuid.UUID) {

}

func (h HttpServer) UpdateReport(w http.ResponseWriter, r *http.Request, reportUUID uuid.UUID) {

}
