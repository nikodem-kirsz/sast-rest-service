package report

import (
	"errors"
	"time"
)

type Report struct {
	UUID          string
	Name          string
	Description   string
	Time          time.Time
	ReportContent string
}

func NewReport(uuid string, name string, description string, time time.Time, reportContent string) (*Report, error) {
	if uuid == "" {
		return nil, errors.New("empty report uuid")
	}
	if name == "" {
		return nil, errors.New("empty name")
	}
	if description == "" {
		return nil, errors.New("empty description")
	}
	if time.IsZero() {
		return nil, errors.New("zero report time")
	}
	if reportContent == "" {
		return nil, errors.New("empty report content")
	}

	return &Report{
		UUID:          uuid,
		Name:          name,
		Description:   description,
		Time:          time,
		ReportContent: reportContent,
	}, nil
}

func UnmarshalReportFromDatabase(
	uuid string,
	name string,
	description string,
	time time.Time,
	reportContent string,
) (*Report, error) {
	re, err := NewReport(uuid, name, description, time, reportContent)
	if err != nil {
		return nil, err
	}

	re.Description = description
	re.Name = name
	re.Time = time
	re.ReportContent = reportContent

	return re, nil
}
