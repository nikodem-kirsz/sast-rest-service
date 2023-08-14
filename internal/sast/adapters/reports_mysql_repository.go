package adapters

import (
	"context"
	"errors"
	"time"

	"github.com/nikodem-kirsz/sast-rest-service/internal/sast/app/query"
	"github.com/nikodem-kirsz/sast-rest-service/internal/sast/domain/report"
	"gorm.io/gorm"
)

type mysqlReportModel struct {
	ID            int       `gorm:"primary_key"`
	UUID          string    `gorm:"column:uuid"`
	Name          string    `gorm:"column:name"`
	Description   string    `gorm:"column:description"`
	Time          time.Time `gorm:"column:time"`
	ReportContent string    `gorm:"column:report_content"`
}

type MySQLRepository struct {
	db *gorm.DB
}

func NewMySQLRepository(db *gorm.DB) *MySQLRepository {
	if db == nil {
		panic("missing db")
	}

	return &MySQLRepository{db: db}
}

func (r MySQLRepository) CreateReport(ctx context.Context, re *report.Report) error {
	model := mysqlReportModel{
		UUID:          re.UUID(),
		Name:          re.Name(),
		Description:   re.Description(),
		Time:          re.Time(),
		ReportContent: re.ReportContent(),
	}

	if err := r.db.Create(&model).Error; err != nil {
		return err
	}

	return nil
}

func (r MySQLRepository) GetReport(ctx context.Context, reportUUID string) (query.Report, error) {
	var model mysqlReportModel
	if err := r.db.Where("uuid = ?", reportUUID).First(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return query.Report{}, nil
		}
		return query.Report{}, err
	}

	return query.Report{
		UUID:          model.UUID,
		Name:          model.Name,
		Description:   model.Description,
		Time:          model.Time,
		ReportContent: model.ReportContent,
	}, nil
}

func (r MySQLRepository) DeleteReport(ctx context.Context, reportUUID string) error {
	if err := r.db.Where("uuid = ?", reportUUID).Delete(&mysqlReportModel{}).Error; err != nil {
		return err
	}
	return nil
}

func (r MySQLRepository) UpdateReport(ctx context.Context, updatedReport *report.Report) error {
	model := mysqlReportModel{
		UUID:          updatedReport.UUID(),
		Name:          updatedReport.Name(),
		Description:   updatedReport.Description(),
		Time:          updatedReport.Time(),
		ReportContent: updatedReport.ReportContent(),
	}

	if err := r.db.Model(&mysqlReportModel{}).Where("uuid = ?", updatedReport.UUID).Updates(&model).Error; err != nil {
		return err
	}

	return nil
}

func (r MySQLRepository) GetAllReports(ctx context.Context) ([]query.Report, error) {
	var models []mysqlReportModel
	if err := r.db.Find(&models).Error; err != nil {
		return nil, err
	}

	var reports []query.Report
	for _, model := range models {
		reports = append(reports, query.Report{
			UUID:          model.UUID,
			Name:          model.Name,
			Description:   model.Description,
			Time:          model.Time,
			ReportContent: model.ReportContent,
		})
	}

	return reports, nil
}
