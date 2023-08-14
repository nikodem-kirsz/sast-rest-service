package query

import (
	"time"
)

type Report struct {
	UUID          string
	Name          string
	Description   string
	Time          time.Time
	ReportContent string
}
