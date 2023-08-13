package service

import (
	"context"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/nikodem-kirsz/sast-rest-service/internal/sast/adapters"
	"github.com/nikodem-kirsz/sast-rest-service/internal/sast/app"
)

func NewApplication(ctx context.Context) *app.Application {
	client, err := firestore.NewClient(ctx, os.Getenv("GCP_PROJECT"))
	if err != nil {
		panic(err)
	}

	reportsRepository := adapters.NewReportsFireStoreRepository(client)

	// logger := logrus.NewEntry(logrus.StandardLogger())
	// metricsClient := metrics.noOp{}

	return app.NewApplication(reportsRepository)
}