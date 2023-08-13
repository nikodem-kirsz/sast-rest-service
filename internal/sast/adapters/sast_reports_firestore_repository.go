package adapters

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/nikodem-kirsz/sast-rest-service/internal/sast/app/query"
	"github.com/nikodem-kirsz/sast-rest-service/internal/sast/domain/report"
	"github.com/pkg/errors"
	"google.golang.org/api/iterator"
)

type ReportModel struct {
	UUID          string    `firestore:"Uuid"`
	Name          string    `firestore:"Name"`
	Description   string    `firestore:"Description"`
	Time          time.Time `firestore:"Time"`
	ReportContent string    `firestore:"ReportContent"`
}

type ReportsFireStoreRepository struct {
	firestoreClient *firestore.Client
}

func NewReportsFireStoreRepository(
	firestoreClient *firestore.Client,
) ReportsFireStoreRepository {
	return ReportsFireStoreRepository{
		firestoreClient: firestoreClient,
	}
}

func (r ReportsFireStoreRepository) reportsCollection() *firestore.CollectionRef {
	return r.firestoreClient.Collection("reports")
}

func (r ReportsFireStoreRepository) AddReport(ctx context.Context, re *report.Report) error {
	collection := r.reportsCollection()

	reportModel := r.marshalReport(re)

	return r.firestoreClient.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		return tx.Create(collection.Doc(reportModel.UUID), reportModel)
	})
}

func (r ReportsFireStoreRepository) GetReport(ctx context.Context, reportUUID string) (*report.Report, error) {
	firestoreReport, err := r.reportsCollection().Doc(reportUUID).Get(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "unable to get actual docs")
	}

	re, err := r.unmarshalReport(firestoreReport)
	if err != nil {
		return nil, err
	}

	return re, nil
}

func (r ReportsFireStoreRepository) DeleteReport(ctx context.Context, reportUUID string) error {
	firestoreReport, err := r.reportsCollection().Doc(reportUUID).Get(ctx)

	if err != nil {
		return errors.Wrap(err, "unable to get actual docs")
	}

	return r.firestoreClient.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		return tx.Delete(firestoreReport.Ref)
	})

}

func (r ReportsFireStoreRepository) UpdateReport(
	ctx context.Context,
	reportUUID string,
	updateFn func(ctx context.Context, re *report.Report) (*report.Report, error),
) error {
	reportsCollection := r.reportsCollection()

	return r.firestoreClient.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		documentRef := reportsCollection.Doc(reportUUID)

		firestoreReport, err := tx.Get(documentRef)
		if err != nil {
			return errors.Wrap(err, "unable to get actual docs")
		}

		tr, err := r.unmarshalReport(firestoreReport)
		if err != nil {
			return err
		}

		updatedReport, err := updateFn(ctx, tr)
		if err != nil {
			return err
		}

		return tx.Set(documentRef, r.marshalReport(updatedReport))
	})
}

func (r ReportsFireStoreRepository) marshalReport(re *report.Report) ReportModel {
	reportModel := ReportModel{
		UUID:          re.UUID(),
		Name:          re.Name(),
		Description:   re.Description(),
		Time:          re.Time(),
		ReportContent: re.ReportContent(),
	}

	return reportModel
}

func (r ReportsFireStoreRepository) unmarshalReport(doc *firestore.DocumentSnapshot) (*report.Report, error) {
	reportModel := ReportModel{}
	err := doc.DataTo(&reportModel)
	if err != nil {
		return nil, errors.Wrap(err, "unable to load document")
	}

	return report.UnmarshalReportFromDatabase(
		reportModel.UUID,
		reportModel.Name,
		reportModel.Description,
		reportModel.Time,
		reportModel.ReportContent,
	)
}

func (r ReportsFireStoreRepository) GetAllReports(ctx context.Context) ([]query.Report, error) {
	query := r.
		reportsCollection().
		Query

	iter := query.Documents(ctx)

	return r.reportModelsToQuery(iter)
}

func (r ReportsFireStoreRepository) reportModelsToQuery(iter *firestore.DocumentIterator) ([]query.Report, error) {
	var reports []query.Report

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		tr, err := r.unmarshalReport(doc)
		if err != nil {
			return nil, err
		}

		queryReport := query.Report{
			UUID:          tr.UUID(),
			Name:          tr.Name(),
			Description:   tr.Description(),
			Time:          tr.Time(),
			ReportContent: tr.ReportContent(),
		}

		reports = append(reports, queryReport)
	}

	return reports, nil
}