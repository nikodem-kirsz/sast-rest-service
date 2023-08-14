package app

import (
	"github.com/nikodem-kirsz/sast-rest-service/internal/sast/app/command"
	"github.com/nikodem-kirsz/sast-rest-service/internal/sast/app/query"
)

type Application struct {
	Queries  Queries
	Commands Commands
}

type Queries struct {
	AllReports query.AllReportsHandler
	GetReport  query.GetReportHandler
}

type Commands struct {
	CreateReport command.CreateReportHandler
	DeleteReport command.DeleteReportHandler
	UpdateReport command.UpdateReportHandler
}
