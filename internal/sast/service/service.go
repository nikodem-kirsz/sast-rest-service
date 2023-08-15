package service

import (
	"context"
	"fmt"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/nikodem-kirsz/sast-rest-service/internal/sast/adapters"
	"github.com/nikodem-kirsz/sast-rest-service/internal/sast/app"
	"github.com/nikodem-kirsz/sast-rest-service/internal/sast/app/command"
	"github.com/nikodem-kirsz/sast-rest-service/internal/sast/app/query"
	"gorm.io/gorm"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
)

type MySqlConfig struct {
	Host     string
	User     string
	Password string
	Port     string
	Database string
}

func NewApplication(ctx context.Context) app.Application {
	var reportsRepository adapters.Repository
	switch db := os.Getenv("DB"); db {
	case "FIRESTORE":
		client, err := firestore.NewClient(ctx, os.Getenv("GCP_PROJECT"))
		if err != nil {
			panic(err)
		}
		reportsRepository = adapters.NewReportsFireStoreRepository(client)
	case "MYSQL":
		mysql := connectToMySQL()
		reportsRepository = adapters.NewMySQLRepository(mysql)
	default:
		panic("No valid database provided, check configuration in .env")
	}

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

func connectToMySQL() *gorm.DB {
	dbConfig := MySqlConfig{
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Database: os.Getenv("MYSQL_DATABASE"),
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database,
	)
	// dsn := "root:password@tcp(mysql:3306)/sast_database?charset=utf8mb4&parseTime=True&loc=Local"

	maxRetries := 2
	retryInterval := 5 * time.Second

	var db *gorm.DB
	var err error

	// Initially waiting 2 seconds before attempting connection to mysql to let the database set
	time.Sleep(2 * time.Second)
	// Perform 2 maximum additional retries while couldn't establish connection(Its enough in this case)
	for retries := 0; retries < maxRetries; retries++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Printf("Failed to connect to database (retry %d/%d): %v\n", retries+1, maxRetries, err)
			time.Sleep(retryInterval)
			continue
		}
		break
	}

	if db == nil {
		panic("Failed to connect to database after retries")
	}

	return db
}
