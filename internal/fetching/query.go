package fetching

import (
	"context"

	"github.com/rfdez/diade/kit/errors"
	"github.com/rfdez/diade/kit/query"
)

const (
	CelebrationByDateQueryType query.Type = "query.fetcher.celeration_by_date"
)

// CelebrationByDateQuery is the query to fetch the celebrations for a given date.
type CelebrationByDateQuery struct {
	date string
}

// NewCelebrationByDateQuery creates a new CelebrationByDateQuery.
func NewCelebrationByDateQuery(date string) CelebrationByDateQuery {
	return CelebrationByDateQuery{date: date}
}

// Type returns the type of the query.
func (q CelebrationByDateQuery) Type() query.Type {
	return CelebrationByDateQueryType
}

// Date returns the date of the query.
func (q CelebrationByDateQuery) Date() string {
	return q.date
}

// CelebrationByDateQueryHandler is the handler of the CelebrationByDateQuery.
type CelebrationByDateQueryHandler struct {
	service Service
}

// NewCelebrationByDateQueryHandler creates a new CelebrationByDateQueryHandler.
func NewCelebrationByDateQueryHandler(service Service) CelebrationByDateQueryHandler {
	return CelebrationByDateQueryHandler{service: service}
}

// Handle handles the CelebrationByDateQuery.
func (h CelebrationByDateQueryHandler) Handle(ctx context.Context, query query.Query) (query.Response, error) {
	q, ok := query.(CelebrationByDateQuery)
	if !ok {
		return nil, errors.New("invalid query type")
	}

	return h.service.CelebrationsByDate(ctx, q.Date())
}
