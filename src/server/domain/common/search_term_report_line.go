package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionSearchTermReportLine = "search_term_report_lines"
)

// SearchTermReportLine represents a search term record (for statistics)
type SearchTermReportLine struct {
	ID      bson.ObjectID `bson:"_id,omitempty"`
	Keyword string        `bson:"keyword"`
	Count   int           `bson:"count"`
}

type SearchTermReportLineRepository interface {
	CreateMany(c context.Context, items []SearchTermReportLine) error
	Create(c context.Context, search_term_record_line *SearchTermReportLine) error
	Update(c context.Context, search_term_record_line *SearchTermReportLine) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]SearchTermReportLine, error)
	FetchByID(c context.Context, ID string) (SearchTermReportLine, error)
}

type SearchTermReportLineUsecase interface {
	CreateMany(c context.Context, items []SearchTermReportLine) error
	FetchByID(c context.Context, ID string) (SearchTermReportLine, error)
	Create(c context.Context, search_term_record_line *SearchTermReportLine) error
	Update(c context.Context, search_term_record_line *SearchTermReportLine) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]SearchTermReportLine, error)
}
