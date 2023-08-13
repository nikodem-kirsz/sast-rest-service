package report

import (
	"errors"
	"time"
)

type Report struct {
	uuid          string
	name          string
	description   string
	time          time.Time
	reportContent string
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
		uuid:          uuid,
		name:          name,
		description:   description,
		time:          time,
		reportContent: reportContent,
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

	re.description = description
	re.name = name
	re.time = time
	re.reportContent = reportContent

	return re, nil
}

func (r Report) UUID() string {
	return r.uuid
}

func (r Report) Name() string {
	return r.name
}

func (r Report) Description() string {
	return r.description
}

func (r Report) Time() time.Time {
	return r.time
}

func (r Report) ReportContent() string {
	return r.reportContent
}
