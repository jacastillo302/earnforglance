package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionSearchTermReportLine = "search_term_report_lines"
)

// SearchTermReportLine represents a search term record (for statistics)
type SearchTermReportLine struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Keyword string             `bson:"keyword"`
	Count   int                `bson:"count"`
}
