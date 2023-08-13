package service

import (
	"context"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/nikodem-kirsz/sast-rest-service/internal/sast/adapters"
	"github.com/nikodem-kirsz/sast-rest-service/internal/sast/app"
	"github.com/nikodem-kirsz/sast-rest-service/internal/sast/app/query"
	"github.com/sirupsen/logrus"
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
		},
	}
}
