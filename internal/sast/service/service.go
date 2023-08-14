package service

import (
	"context"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/nikodem-kirsz/sast-rest-service/internal/sast/adapters"
	"github.com/nikodem-kirsz/sast-rest-service/internal/sast/app"
	"github.com/nikodem-kirsz/sast-rest-service/internal/sast/app/command"
	"github.com/nikodem-kirsz/sast-rest-service/internal/sast/app/query"
	"gorm.io/gorm"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
)

func NewApplication(ctx context.Context) app.Application {
	client, err := firestore.NewClient(ctx, os.Getenv("GCP_PROJECT"))
	if err != nil {
		panic(err)
	}

	reportsRepository := adapters.NewReportsFireStoreRepository(client)

	logger := logrus.NewEntry(logrus.StandardLogger())

	return app.Application{
		Queries: app.Queries{
			AllReports: query.NewAllReportsHandler(reportsRepository, logger),
			GetReport:  query.NewGetReportHandler(reportsRepository, logger),
		},
		Commands: app.Commands{
			CreateReport: command.NewCreateReportHandler(reportsRepository, logger),
			UpdateReport: command.NewUpdateReportHandler(reportsRepository, logger),
			DeleteReport: command.NewDeleteReportHandler(reportsRepository, logger),
		},
	}
}

func NewApplicationWithSQL(ctx context.Context) app.Application {
	dsn := "root:password@tcp(127.0.0.1:3306)/sast_database?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	reportsRepository := adapters.NewMySQLRepository(db)

	logger := logrus.NewEntry(logrus.StandardLogger())

	return app.Application{
		Queries: app.Queries{
			AllReports: query.NewAllReportsHandler(reportsRepository, logger),
			GetReport:  query.NewGetReportHandler(reportsRepository, logger),
		},
		Commands: app.Commands{
			CreateReport: command.NewCreateReportHandler(reportsRepository, logger),
			UpdateReport: command.NewUpdateReportHandler(reportsRepository, logger),
			DeleteReport: command.NewDeleteReportHandler(reportsRepository, logger),
		},
	}
}
